/**
 * Return (x days ago) from given timestamp
 * @param {*} timestamp 
 * @returns 
 */
const relativeDate = (timestamp) => {
    const rtf = new Intl.RelativeTimeFormat('en', {
      numeric: 'auto',
    });
    const oneDayInMs = 1000 * 60 * 60 * 24;
    const daysDifference = Math.round(
      (timestamp - new Date().getTime()) / oneDayInMs,
    );
  
    return rtf.format(daysDifference, 'day');
}
 
/**
 * Converts currency sign to currency symbol
 * @param {*} currencyName 
 * @returns 
 */
const currencySign = currencyName => ({
  "USD": "$"
}[currencyName] || currency)