const HOSTNAME = "192.168.193.208:8082"

const renderTransaction = (transactionsHTML, transaction) => {
    return `${transactionsHTML}<div class="transaction">${transaction.Value}</div>`
}

const paint = (transactions, transactionsHtml) => {
    document.getElementById(transactions).innerHTML = transactionsHtml
}

fetch(`http://${HOSTNAME}/transactions`)
    .then(req => req.json())
    .then(res => paint(
        "transactions",
        res.transactions.reduce(renderTransaction, '')
    ))