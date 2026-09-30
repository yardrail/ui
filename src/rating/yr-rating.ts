// <yr-rating value="4.5" reviews="128"> — a 5-star rating row, rounded to
// the nearest whole star (consolidates depot's numeric detail-rating and
// review-card's integer filled/empty star loop into one widget). Renders
// its own light-DOM children from attributes (textContent only).
export class YrRating extends HTMLElement {
  static observedAttributes = ['value', 'reviews'];

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  private render() {
    const raw = parseFloat(this.getAttribute('value') ?? '0');
    const filled = Math.max(0, Math.min(5, Math.round(raw)));
    const reviews = this.getAttribute('reviews');

    this.replaceChildren();

    const stars = document.createElement('span');
    stars.className = 'yr-rating-stars';
    stars.textContent = '★'.repeat(filled) + '☆'.repeat(5 - filled);
    this.append(stars);

    if (reviews) {
      const count = document.createElement('span');
      count.className = 'yr-rating-count';
      count.textContent = `(${reviews} reviews)`;
      this.append(count);
    }
  }
}

customElements.define('yr-rating', YrRating);
