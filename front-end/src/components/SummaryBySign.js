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
  const [income, outgoing] = props.payments.reduce((acc, cur) => {
    if (cur.price > 0)
      acc[0] += cur.price;
    else
      acc[1] += -cur.price;
    return acc;
  }, [0, 0]);

  return (
    <>
      <h6 className={props.className}><b style={{'color': '#02d505'}}>수입: </b>{income.toLocaleString()}원</h6>
      <h6 className={props.className}><b style={{'color': '#fd2926'}}>지출: </b>{outgoing.toLocaleString()}원</h6>
    </>
  );
}

export default SummaryBySign;
