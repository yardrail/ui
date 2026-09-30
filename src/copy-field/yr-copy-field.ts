// <yr-copy-field value="tok_live_..."> renders a code display + copy
// button. Clicking copies `value` to the clipboard and shows "Copied!"
// feedback for 1.5s. Replaces depot's .code-block + .deployment-token-copy
// pattern (a code block wired to a separately-authored copy button).
export class YrCopyField extends HTMLElement {
  static observedAttributes = ['value'];
  private resetTimer: number | undefined;

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  disconnectedCallback() {
    if (this.resetTimer !== undefined) window.clearTimeout(this.resetTimer);
  }

  private render() {
    const value = this.getAttribute('value') ?? '';
    this.replaceChildren();

    const code = document.createElement('code');
    code.className = 'yr-copy-field-code';
    code.textContent = value;

    const button = document.createElement('button');
    button.type = 'button';
    button.className = 'yr-copy-field-button';
    button.textContent = 'Copy';
    button.addEventListener('click', () => this.handleCopy(value, button));

    this.append(code, button);
  }

  private async handleCopy(value: string, button: HTMLButtonElement) {
    try {
      await navigator.clipboard.writeText(value);
      button.textContent = 'Copied!';
    } catch {
      button.textContent = 'Copy failed';
    }
    if (this.resetTimer !== undefined) window.clearTimeout(this.resetTimer);
    this.resetTimer = window.setTimeout(() => {
      button.textContent = 'Copy';
    }, 1500);
  }
}

customElements.define('yr-copy-field', YrCopyField);
