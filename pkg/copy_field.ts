// <yr-copy-field> — code + copy button structure is server-rendered by the
// templ component. This class only wires up the clipboard copy behavior.
export class YrCopyField extends HTMLElement {
  private resetTimer: number | undefined;

  connectedCallback() {
    const button = this.querySelector('.yr-copy-field-button');
    if (button) {
      button.addEventListener('click', () => {
        const value = (button as HTMLElement).dataset.value ?? '';
        this.handleCopy(value, button as HTMLButtonElement);
      });
    }
  }

  disconnectedCallback() {
    if (this.resetTimer !== undefined) window.clearTimeout(this.resetTimer);
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
