// <yr-live-duration started-at="2026-06-07T09:00:00Z"> renders an HH:MM:SS
// counter that ticks every second client-side. Used on cards/lists where the
// duration since a start time should stay live without round-tripping to the
// server. Relocated from yardrail-ui's private live-duration element.
export class YrLiveDuration extends HTMLElement {
  static observedAttributes = ['started-at'];
  private intervalId: number | undefined;

  connectedCallback() {
    this.tick();
    this.intervalId = window.setInterval(() => this.tick(), 1000);
  }

  disconnectedCallback() {
    if (this.intervalId !== undefined) window.clearInterval(this.intervalId);
  }

  attributeChangedCallback() {
    this.tick();
  }

  private tick() {
    const startedAt = this.getAttribute('started-at');
    if (!startedAt) {
      this.textContent = '--:--:--';
      return;
    }
    const elapsedMs = Math.max(0, Date.now() - new Date(startedAt).getTime());
    this.textContent = formatDuration(elapsedMs);
  }
}

function formatDuration(ms: number): string {
  const totalSeconds = Math.floor(ms / 1000);
  const h = Math.floor(totalSeconds / 3600);
  const m = Math.floor((totalSeconds % 3600) / 60);
  const s = totalSeconds % 60;
  return [h, m, s].map((n) => String(n).padStart(2, '0')).join(':');
}

customElements.define('yr-live-duration', YrLiveDuration);
