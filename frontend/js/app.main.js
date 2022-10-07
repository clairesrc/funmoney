window.env = "prod"

/**
 * ID attribute of the HTML node to paint the app to.
 */
const WRAPPER_ID = "main"

/**
 * Global app state.
 */
let appState = {
    sections: {},
    listeners: [],
    data: {}
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
const addSection = (title, name, renderer) => appState.sections[name] = { title, name, renderer, type: "section" }

/**
 * Add form to list of sections to fetch and render.
 */
const addFormSection = (title, name, renderer, onSubmit) => {
    appState.sections[name] = { title, name, renderer, type: "form" }
    appState.listeners.push(() => {
        cl("adding event listener for " + name)
        document.getElementById(name).addEventListener("submit", onSubmit)
    })
}

/**
 * Render all sections.
 */
const renderApp = (appData, sectionData, sections, listeners) => {
    wrapNode = document.getElementById(WRAPPER_ID)
    wrapNode.innerHTML = sectionData.reduce(
        (content, { name, data }) => {
            // Render section
            cl(`rendering ${name} section`)
            const section = sections[name]
            if (section.type == "section") {
                ({ title, renderer } = section)
                return content + renderSection(name, title, renderer(appData, data))
            } else if (section.type == "form") {
                ({ title, renderer, onSubmit } = section)
                return content + renderSection(name, title, renderer(appData))
            }
        },
        ''
    )
    listeners.forEach(listener => listener())
}

/**
 * GET data from API endpoint.
 * Assumes section name = endpoint = response top-level key.
 * Appends a timestamp cache bust.
 */
const getData = async name => fetch(`http://${getHostname()}/${name}${'?' + Date.now()}`)
    .then(res => res.json())
    .then(data => {
        // Add data to global dictionary
        appState.data[name] = data[name]

        return { name, "data": data[name] }
    })

/**
 * Get app data, then get section data, then render sections. 
 */
const main = ({ sections, listeners }) => {
    renderApp({}, Object.keys(sections).map(section => ({ name: section, data: {} })), sections, listeners)
    getData('app')
        .then(({ data }) =>
            Promise.all(Object.keys(sections)
                .map(key => sections[key].type == "section" ? getData(key) : { name: key, data: {} })
            ).then(sectionData => renderApp(data, sectionData, sections, listeners))
                .catch(console.error)
        )
}

/** 
 * Add sections.
 */
addSection("My Balance", "balance", balanceRenderer)
addFormSection("Add Transaction", "addTransaction", addTransactionRenderer, addTransactionOnSubmit)
addSection("Recent Transactions", "transactions", transactionRenderer)

/**
 * Run app.
 */
main(appState)