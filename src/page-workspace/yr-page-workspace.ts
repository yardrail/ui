export class YrPageWorkspace extends HTMLElement {
  connectedCallback() {
    this.setAttribute('role', 'main');
  }
}

customElements.define('yr-page-workspace', YrPageWorkspace);
