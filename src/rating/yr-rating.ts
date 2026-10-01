// <yr-rating value="4.5" reviews="128"> — a 5-star rating row, rounded to
// the nearest whole star (consolidates depot's numeric detail-rating and
// review-card's integer filled/empty star loop into one widget). Renders
// Lucide star SVGs: filled stars use fill="currentColor", empty stars are
// stroke-only.
const STAR_PATH =
  'M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z';

function starSVG(filled: boolean): SVGElement {
  const ns = 'http://www.w3.org/2000/svg';
  const svg = document.createElementNS(ns, 'svg');
  svg.setAttribute('width', '16');
  svg.setAttribute('height', '16');
  svg.setAttribute('viewBox', '0 0 24 24');
  svg.setAttribute('fill', filled ? 'currentColor' : 'none');
  svg.setAttribute('stroke', 'currentColor');
  svg.setAttribute('stroke-width', '2');
  svg.setAttribute('stroke-linecap', 'round');
  svg.setAttribute('stroke-linejoin', 'round');
  const path = document.createElementNS(ns, 'path');
  path.setAttribute('d', STAR_PATH);
  svg.appendChild(path);
  return svg;
}

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
    for (let i = 0; i < 5; i++) {
      stars.appendChild(starSVG(i < filled));
    }
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
