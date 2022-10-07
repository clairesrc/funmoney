const addTransactionRenderer = (appData) => `
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
    alert('hi :)');
    return false
}