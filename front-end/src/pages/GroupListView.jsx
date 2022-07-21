import '@/pages/GroupListView.scss';
import { AiOutlinePlus } from 'react-icons/ai';
import { Button } from 'react-bootstrap';
import useRequest from '@/common/useRequest';
import ProcessingSpinner from '@/common/ProcessingSpinner';
import SummaryBySign from '@/components/SummaryBySign';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

function _buildDateRange() {
  const today = new Date();
  return {
    dateFrom: `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}`,
    dateTo: `${today.getFullYear()}-${String(today.getMonth() + 2).padStart(2, '0')}`
  }
}

function GroupItemView(props) {
  const navigate = useNavigate();

  const { group, setProcessing } = props;
  const [_processing, payments] = useRequest(
    '/v1/payment', {
      groupId: group.id,
      ..._buildDateRange()
    }, [], []);

  useEffect(() => {
    setProcessing(_processing);
  }, [_processing, setProcessing]);

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
      </main>
      <footer className="group-item-footer">
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          disabled
        >내역 등록</Button>
        <div className="action-divider" />
        <Button
          size="sm"
          variant="clear"
          className="footer-action"
          onClick={() => {
            navigate(`/payments/${group.id}`);
          }}
        >내역 확인</Button>
      </footer>
    </section>
  );
}

export default function GroupListView(props) {
  const [processing, setProcessing] = useState();
  const [_processing, groups] = useRequest(
    '/v1/group', { email: props.userInfo.email }, [], []);

  useEffect(() => {
    setProcessing(_processing);
  }, [_processing]);

  return (
    <article className="abs-group">
      {groups.map(group =>
        <GroupItemView
          group={group}
          key={group.id}
          setProcessing={setProcessing}
        />)}
      <section className="group-creation">
        <Button
          className="w-100"
          variant="outline-primary"
          disabled
        ><AiOutlinePlus /></Button>
      </section>
      <ProcessingSpinner processing={processing} />
    </article>
  );
}
