
/**
 * ID attribute of the HTML node to paint the app to.
 */
const WRAPPER_ID = "main"

/**
 * Global app state.
 */
let appState = {
    sections: {},
}

/**
 * Override hostname. Useful for remote development.
 */
 const getHostname = () => {
    const urlParams = new URLSearchParams(window.location.search)
    const hostname = urlParams.get('hostname')
    return hostname ? hostname : "localhost:8082"
}

/**
 * Render Balance section content.
 */
const balanceRenderer = (appData, balance) => `
    <p>
        Current balance: <span class="balance-figure">${balance ? balance : 0.00}</span>
    </p>`

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
`, '') : `Error: ${transactions}`

/**
 * Render section with content.
 */
const renderSection = (name, title, sectionContent) => `
    <section id="${name}">
        <h1>${title}</h1>
        <div class="content">
            ${sectionContent}
        </div>
    </section>
    `

/**
 * Add section to list of sections to fetch and render.
 */
const addSection = (title, name, renderer) => appState.sections[name] = {title, name, renderer}

/**
 * Render all sections.
 */
const renderApp = (appData, sectionData, sections) => {
    wrapNode = document.getElementById(WRAPPER_ID)
    wrapNode.innerHTML = sectionData.reduce(
        (content, { name, data }) => {
            ({ title, renderer } = sections[name])
            return content + renderSection(name, title, renderer(appData, data))
        },
        ''
    )
}

/**
 * GET data from API endpoint.
 * Assumes section name = endpoint = response top-level key
 */
const getData = async name => fetch(`http://${getHostname()}/${name}`)
    .then(res => res.json())
    .then(data => {
        return { name, "data": data[name] }
    })

/**
 * Get app data, then get section data, then render sections. 
 */
const main = appState =>
getData('app')
.then(appData => 
    Promise.all(Object.keys(appState.sections)
        .map(key => getData(appState.sections[key].name))
    ).then(data => renderApp(appData.data, data, appState.sections))
        .catch(console.error)
)

/** 
 * Add sections.
 */
addSection("My Balance", "balance", balanceRenderer)
addSection("Recent Transactions", "transactions", transactionRenderer)

/**
 * Run app.
 */
main(appState)