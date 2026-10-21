// <yr-journal-log>...</yr-journal-log> — a scrollable container for
// workflow run journal entries. Sets role="log" and aria-label automatically
// so screen readers announce new entries. Light DOM, styled by
// yr-journal-log.css.
export class YrJournalLog extends HTMLElement {
  connectedCallback() {
    if (!this.getAttribute('role')) this.setAttribute('role', 'log');
    if (!this.getAttribute('aria-label')) this.setAttribute('aria-label', 'Run journal');
  }
}

customElements.define('yr-journal-log', YrJournalLog);
