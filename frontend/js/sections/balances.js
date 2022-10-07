/**
 * Render Balance section content.
 */
const balanceRenderer = (appData, balance) => appData && appData.currency && appData.cap ? `
    <p>
        Current balance: <span class="balance-currency">${currencySign(appData.currency)}</span><span class="balance-value">${balance ? balance : 0.00}</span><span class="balance-out-of">/</span><span class="balance-currency">${currencySign(appData.currency)}</span><span class="balance-value">${appData.cap ? appData.cap : 0.00}</span>
    </p>` : 'Counting beans&hellip;'