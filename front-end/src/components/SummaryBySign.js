import './SummaryBySign.css';
import {useEffect, useReducer} from 'react';
import * as Api from '../common/Api';
import ProcessingSpinner from '../common/ProcessingSpinner';

/**
 * @typedef {Object} Payment
 * @property {string} category
 * @property {string} date
 * @property {string} method
 * @property {number} monthlyInstallment
 * @property {string} name
 * @property {number} price
 */

function SummaryBySign(props) {
  const [state, dispatch] = useReducer((_, action) => {
    const /** @type {Array.<Payment>} */ payments = action.data ?? [];
    return payments.reduce((acc, cur) => {
      if (cur.price > 0)
        acc.income += cur.price;
      else
        acc.outgoing += -cur.price;
      return acc;
    }, { processing: action.processing, income: 0, outgoing: 0 });
  }, { processing: false, income: 0, outgoing: 0 });

  useEffect(() => {
    fetchData().then();

    async function fetchData() {
      const monthDateStr = ``+
        `${new Date().getFullYear()}-${String(props.currentMonth + 1).padStart(2, '0')}`;

      await Api.getWithDispatch(dispatch, '/payments', {
        start: monthDateStr,
        end: monthDateStr
      });
    }
  }, [props.currentMonth]);

  return (
    <article
      className="abs-summary-by-sign"
      style={{'marginTop': `${props.mt}em`}}
    >
      <h6><b style={{'color': '#02d505'}}>수입: </b>{state.income.toLocaleString()}원</h6>
      <h6><b style={{'color': '#fd2926'}}>지출: </b>{state.outgoing.toLocaleString()}원</h6>
      <ProcessingSpinner processing={state.processing}/>
    </article>
  );
}

export default SummaryBySign;
