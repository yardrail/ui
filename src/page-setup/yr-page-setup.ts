export class YrPageSetup extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'main');
  }
}

export class YrPageSetupContent extends HTMLElement {}

customElements.define('yr-page-setup', YrPageSetup);
customElements.define('yr-page-setup-content', YrPageSetupContent);
