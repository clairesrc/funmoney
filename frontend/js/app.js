const HOSTNAME = "192.168.193.208:8082"
const WRAPPER_ID = "main"

let appState = {
    sections: {}
}

const balanceRenderer = balance => `
    <p>
        Current balance: <span class="balance-figure">${balance ? balance : 0.00}</span>
    </p>`

const transactionRenderer = transactions => transactions ? transactions.reduce((transactionsHTML, { Value, Timestamp, Comment, ID }) =>  `${transactionsHTML}
<div class="transaction" data-transaction-id="${ID}">
    <div class="amount-wrap">
        <span class="amount-currency">$</span>
        <span class="amount-value">${Value ? Value : 0.00}</span>
    </div>
    <span class="timestamp">${Timestamp ? relativeDate(new Date(Timestamp * 1000)) : 0.00}</span>
    <div class="comment">${Comment}</div>
</div>
`, '') : `Error: ${transactions}`

const renderSection = (name, title , sectionContent) => `
    <section id="${name}">
        <h1>${title}</h1>
        <div class="content">
            ${sectionContent}
        </div>
    </section>
    `

const addSection = (title, name, renderer) => appState.sections[name] = {title, name, renderer}

const renderApp = (appData, sections) => {
    wrapNode = document.getElementById(WRAPPER_ID)
    wrapNode.innerHTML = appData.reduce(
        (content, { name, data }) => {
            ({ title, renderer } = sections[name])
            return content + renderSection(name, title, renderer(data))
        },
        ''
    )
}

const getSectionData = async name => fetch(`http://${HOSTNAME}/${name}`)
    .then(res => res.json())
    .then(data => { return { name, "data": data[name] } })

const main = appState => Promise.all(
    Object.keys(appState.sections).map(key => getSectionData(appState.sections[key].name))
).then(data => renderApp(data, appState.sections))
.catch(console.error)

addSection("My Balance", "balance", balanceRenderer)
addSection("Recent Transactions", "transactions", transactionRenderer)

main(appState)