const addTransactionRenderer = (appData, added = false) => `
    ${added ? '<div class="message">Added new transaction</div>' : ''}
    <form>
        <label for="addTransaction-value">${appData.currency || ''} Amount</label>
        <input type="text" name="value" value="" id="addTransaction-value" placeholder="Format: 100.00, Required" aria-required="true" required>
        <label for="addTransaction-comment">Comment</label>
        <textarea name="comment" value="" id="addTransaction-comment" placeholder="Optional"></textarea>
        <input type="submit" value="Add" id="addTransaction-submit">
    </form>
    `
const addTransactionOnSubmit = event => {
    event.preventDefault()

    fetch(`http://${getHostname()}/transaction`, {
        method: "post",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            comment: document.getElementById("addTransaction-comment").value,
            value: parseFloat(document.getElementById("addTransaction-value").value)
        })
    }).then((response) => response.json())
        .then((result) => {
            // @TODO: manually updating global state and re-rendering sections is untenable! figure out a cleaner way to handle this.
            // Update global state
            appState.data.transactions = appState.data.transactions ? [result.transaction, ...appState.data.transactions] : [result.transaction]
            appState.data.balance += result.transaction.Value

            // Update sections
            document.querySelector("#transactions .content").innerHTML = transactionRenderer({ currency: appState.data.app.currency }, appState.data.transactions)
            document.querySelector("#balance .content").innerHTML = balanceRenderer(appState.data.app, appState.data.balance)
            document.querySelector("#addTransaction .content").innerHTML = addTransactionRenderer(appState.data.app, true)

        })

    return false
}