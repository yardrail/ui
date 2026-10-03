export class YrSetupCard extends HTMLElement {}

export class YrSetupCardProgress extends HTMLElement {
  static observedAttributes = ['steps', 'current'];

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  private render() {
    const steps = parseInt(this.getAttribute('steps') || '0', 10);
    const current = parseInt(this.getAttribute('current') || '1', 10);
    this.innerHTML = '';
    for (let i = 1; i <= steps; i++) {
      const dot = document.createElement('span');
      dot.className = 'yr-setup-card-dot';
      if (i < current) dot.classList.add('yr-setup-card-dot--done');
      if (i === current) dot.classList.add('yr-setup-card-dot--active');
      this.appendChild(dot);
    }
  }
}

export class YrSetupCardHeader extends HTMLElement {}
export class YrSetupCardBody extends HTMLElement {}

export class YrSetupCardFooter extends HTMLElement {}

customElements.define('yr-setup-card', YrSetupCard);
customElements.define('yr-setup-card-progress', YrSetupCardProgress);
customElements.define('yr-setup-card-header', YrSetupCardHeader);
customElements.define('yr-setup-card-body', YrSetupCardBody);
customElements.define('yr-setup-card-footer', YrSetupCardFooter);
