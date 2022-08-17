import '@/pages/GroupListView.scss';
import { AiOutlinePlus } from 'react-icons/ai';
import { Button } from 'react-bootstrap';
import useRequest from '@/common/useRequest';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import SummaryBySign from '@/components/SummaryBySign';
import { useNavigate } from 'react-router-dom';

function _buildDateRange() {
  const today = new Date();
  return {
    dateFrom: `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}`,
    dateTo: `${today.getFullYear()}-${String(today.getMonth() + 2).padStart(2, '0')}`
  }
}

function GroupItemView({ group }) {
  const navigate = useNavigate();
  const [_processing, payments] = useRequest(
    `/v1/group/${group.id}/payment`,
    _buildDateRange(), [], []);

  return (
    <section
      className="group-item"
      key={group.id}
    >
      <h5>{group.name}</h5>
      <main className="group-item-main">
        <SummaryBySign
          payments={payments}
          className="bilateral-align"
        />
        <ProcessingSpinner processing={_processing} />
      </main>
      <footer className="group-item-footer">
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          onClick={() => {
            navigate('/payments/register', {
              state: { gid: group.id }
            });
          }}
        >내역 등록</Button>
        <div className="action-divider" />
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          onClick={() => {
            navigate(`/payments`, {
              state: {
                gid: group.id,
                payments
              }
            });
          }}
        >내역 확인</Button>
      </footer>
    </section>
  );
}

export default function GroupListView(props) {
  const [_processing, groups] = useRequest(
    '/v1/group', { email: props.userInfo.email }, [], []);

  return (
    <article className="abs-group">
      {groups.map(group =>
        <GroupItemView
          group={group}
          key={group.id}
        />)}
      <section className="group-creation">
        <Button
          className="w-100"
          variant="outline-primary"
          disabled
        ><AiOutlinePlus /></Button>
      </section>
      <ProcessingSpinner processing={_processing} />
    </article>
  );
}
