
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
const main = ({sections}) =>
getData('app')
.then(({data}) => 
    Promise.all(Object.keys(sections)
        .map(key => getData(sections[key].name))
    ).then(sectionData => renderApp(data, sectionData, sections))
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