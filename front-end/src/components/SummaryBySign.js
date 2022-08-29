import React from 'react';

/** @typedef {string} KorStdTimeStr - 2022-07-10T14:39:33.959Z */
/**
 * @typedef {Object} Payment
 * @property {string} category
 * @property {KorStdTimeStr} created_at
 * @property {KorStdTimeStr} date
 * @property {string} groupId
 * @property {string} id
 * @property {string} modUserId
 * @property {number} monthlyInstallment
 * @property {string} name
 * @property {string} paymentMethodId
 * @property {Array} paymentMethods
 * @property {number} price
 * @property {string} regUserId
 * @property {KorStdTimeStr} updated_at
 */

function SummaryBySign({ payments, className }) {
  const [income, outgoing] = payments.reduce((acc, cur) => {
    if (cur.price > 0)
      acc[0] += cur.price;
    else
      acc[1] += -cur.price;
    return acc;
  }, [0, 0]);

  return (
    <>
      <h6 className={className}>
        <b style={{'color': '#02d505'}}>수입: </b>{income.toLocaleString()}원
      </h6>
      <h6 className={className}>
        <b style={{'color': '#fd2926'}}>지출: </b>{outgoing.toLocaleString()}원
      </h6>
    </>
  );
}

export default React.memo(SummaryBySign);
