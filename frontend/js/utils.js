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

const currencySigns = {
  "USD": "$"
}
  
const currencySign = currencyName => currencySigns[currencyName] || currency