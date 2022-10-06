const HOSTNAME = "192.168.193.208:8082"

const balanceRenderer = balance => `
    <p>
        Current balance: <span class="balance-figure">${balance}</span>
    </p>`

const transactionRenderer = transactions => transactions.reduce((transactionsHTML, {Value}) => `${transactionsHTML}<div class="transaction">${Value}</div>`, '')

const paint = (transactions, transactionsHtml) => {
    document.getElementById(transactions).innerHTML = transactionsHtml
}

const addSection = (name, renderer) => fetch(`http://${HOSTNAME}/${name}`)
.then(req => req.json())
.then(res => paint(
    name,
    renderer(res[name]))
)

addSection("transactions", transactionRenderer)
addSection("balance", balanceRenderer)
