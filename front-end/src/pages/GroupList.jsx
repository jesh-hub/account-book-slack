import '@/pages/GroupList.scss';
import { Button } from 'react-bootstrap';
import useRouterNavigateWith from '@/common/useRouterNavigateWith';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import SummaryBySign from '@/components/SummaryBySign';
import { useNavigate } from 'react-router-dom';
import { useGetRequest } from '@/common/Api.js';
import { useMemo } from 'react';
import { getDateMonthStr } from '@/common/DateUtil';

function GroupItemView({ group }) {
  const navigateWith = useRouterNavigateWith();
  const dateParams = useMemo(() => {
    const dateFrom = new Date(),
      dateTo = new Date();
    dateTo.setMonth(dateTo.getMonth() + 1);
    return {
      dateFrom: getDateMonthStr(dateFrom),
      dateTo: getDateMonthStr(dateTo)
    };
  }, []);
  const [payments, processing] = useGetRequest(`/v1/group/${group.id}/payment`, dateParams);

  return (
    <section className="group-item">
      <h5>{group.name}</h5>
      <main className="group-item-main">
        <SummaryBySign
          payments={payments}
          className="bilateral-align"
        />
        <ProcessingSpinner processing={processing} />
      </main>
      <footer className="group-item-footer">
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          onClick={() => {
            navigateWith('/payments/register', { gid: group.id });
          }}
        >내역 등록</Button>
        <div className="action-divider" />
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          onClick={() => {
            navigateWith(`/payments`, {
              gid: group.id,
              payments
            });
          }}
        >내역 확인</Button>
      </footer>
    </section>
  );
}

export default function GroupList({ userInfo }) {
  const navigate = useNavigate();
  const params = useMemo(() => ({ email: userInfo.email }), [userInfo]);
  const [groups, processing] = useGetRequest('/v1/group', params);

  return (
    <article className="abs-groups">
      {
        groups.map(group =>
          <GroupItemView
            group={group}
            key={group.id}
          />)
      }
      <section className="group-creation">
        <Button
          className="w-100"
          variant="outline-primary"
          size="sm"
          onClick={() => {
            navigate('/groups/register');
          }}
        >새 그룹 만들기</Button>
      </section>
      <ProcessingSpinner processing={processing} />
    </article>
  );
}
