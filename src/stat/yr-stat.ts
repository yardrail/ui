// <yr-stat number="1,204" label="Templates installed"> — a number + label
// pair. Renders its own light-DOM children from attributes (textContent
// only, no innerHTML) so callers just pass the two attributes.
export class YrStat extends HTMLElement {
  static observedAttributes = ['number', 'label'];

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  private render() {
    this.replaceChildren();

    const numberEl = document.createElement('span');
    numberEl.className = 'yr-stat-number';
    numberEl.textContent = this.getAttribute('number') ?? '';

    const labelEl = document.createElement('span');
    labelEl.className = 'yr-stat-label';
    labelEl.textContent = this.getAttribute('label') ?? '';

    this.append(numberEl, labelEl);
  }
}

customElements.define('yr-stat', YrStat);
