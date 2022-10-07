/**
 * Render Balance section content.
 */
const balanceRenderer = (appData, balance) => `
    <p>
        Current balance: <span class="balance-figure">${balance ? balance : 0.00}</span>
    </p>`