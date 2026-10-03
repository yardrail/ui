export class YrPageAuth extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'main');
  }
}

customElements.define('yr-page-auth', YrPageAuth);
