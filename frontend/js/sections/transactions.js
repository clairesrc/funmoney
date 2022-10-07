/**
 * Render Transactions section content.
 */
 const transactionRenderer = ({currency}, transactions) => transactions ? transactions.reduce((transactionsHTML, { Value, Timestamp, Comment, ID }) =>  `${transactionsHTML}
 <div class="transaction" data-transaction-id="${ID}">
     <div class="amount-wrap">
         <span class="amount-currency">${currencySign(currency)}</span><span class="amount-value">${Value ? Value : 0.00}</span>
     </div>
     <span class="timestamp">${Timestamp ? relativeDate(new Date(Timestamp * 1000)) : 0.00}</span>
     <div class="comment">${Comment}</div>
 </div>
 `, '') : `<div class="empty" data-value="${transactions}">No transactions yet!</div>`