type Corner = 'bottom-right' | 'bottom-left' | 'top-right' | 'top-left';

const MARGIN = 24;
const SIZE = 56;

export class YrChatFab extends HTMLElement {
  private startX = 0;
  private startY = 0;
  private prevX = 0;
  private prevY = 0;
  private prevTime = 0;
  private velX = 0;
  private velY = 0;
  private isDragging = false;
  private hasMoved = false;

  connectedCallback() {
    this.trigger?.addEventListener('mousedown', this.onPointerDown as EventListener);
    this.trigger?.addEventListener('click', this.onClick);
    this.querySelector('.yr-chat-fab__dismiss')?.addEventListener('click', this.onDismiss);
  }

  disconnectedCallback() {
    this.trigger?.removeEventListener('mousedown', this.onPointerDown as EventListener);
    this.trigger?.removeEventListener('click', this.onClick);
    this.querySelector('.yr-chat-fab__dismiss')?.removeEventListener('click', this.onDismiss);
    document.removeEventListener('mousemove', this.onPointerMove as EventListener);
    document.removeEventListener('mouseup', this.onPointerUp);
  }

  private get trigger(): Element | null {
    return this.querySelector('.yr-chat-fab__trigger');
  }

  private onClick = (e: Event) => {
    if (this.hasMoved) {
      e.preventDefault();
      e.stopPropagation();
      return;
    }
    this.dispatchEvent(new CustomEvent('yr-chat-toggle', { bubbles: true }));
  };

  private onDismiss = (e: Event) => {
    e.stopPropagation();
    this.hidden = true;
    this.dispatchEvent(new CustomEvent('yr-chat-dismiss', { bubbles: true }));
  };

  private onPointerDown = (e: MouseEvent) => {
    this.startX = e.clientX;
    this.startY = e.clientY;
    this.prevX = e.clientX;
    this.prevY = e.clientY;
    this.prevTime = performance.now();
    this.velX = 0;
    this.velY = 0;
    this.isDragging = true;
    this.hasMoved = false;
    document.addEventListener('mousemove', this.onPointerMove as EventListener);
    document.addEventListener('mouseup', this.onPointerUp);
  };

  private onPointerMove = (e: MouseEvent) => {
    if (!this.isDragging) return;
    const dx = e.clientX - this.startX;
    const dy = e.clientY - this.startY;
    if (!this.hasMoved && Math.abs(dx) < 5 && Math.abs(dy) < 5) return;

    if (!this.hasMoved) {
      this.hasMoved = true;
      this.setAttribute('dragging', '');
      this.removeAttribute('corner');
      this.clearPositionProps();
    }

    const now = performance.now();
    const dt = now - this.prevTime;
    if (dt > 0) {
      this.velX = (e.clientX - this.prevX) / dt;
      this.velY = (e.clientY - this.prevY) / dt;
    }
    this.prevX = e.clientX;
    this.prevY = e.clientY;
    this.prevTime = now;

    this.style.left = `${e.clientX - SIZE / 2}px`;
    this.style.top = `${e.clientY - SIZE / 2}px`;
  };

  private onPointerUp = () => {
    document.removeEventListener('mousemove', this.onPointerMove as EventListener);
    document.removeEventListener('mouseup', this.onPointerUp);
    this.isDragging = false;

    if (!this.hasMoved) return;

    const rect = this.getBoundingClientRect();
    const MOMENTUM = 300;
    const cx = rect.left + rect.width / 2 + this.velX * MOMENTUM;
    const cy = rect.top + rect.height / 2 + this.velY * MOMENTUM;
    const midX = window.innerWidth / 2;
    const midY = window.innerHeight / 2;

    const corner: Corner =
      cy < midY
        ? cx < midX ? 'top-left' : 'top-right'
        : cx < midX ? 'bottom-left' : 'bottom-right';

    const target = this.cornerPosition(corner);

    this.removeAttribute('dragging');
    this.style.top = `${target.top}px`;
    this.style.left = `${target.left}px`;

    const onDone = () => {
      this.removeEventListener('transitionend', onDone);
      this.clearPositionProps();
      this.setAttribute('corner', corner);
    };
    this.addEventListener('transitionend', onDone);

    setTimeout(onDone, 450);
  };

  private cornerPosition(corner: Corner): { top: number; left: number } {
    const w = window.innerWidth;
    const h = window.innerHeight;
    switch (corner) {
      case 'top-left':
        return { top: MARGIN, left: MARGIN };
      case 'top-right':
        return { top: MARGIN, left: w - SIZE - MARGIN };
      case 'bottom-left':
        return { top: h - SIZE - MARGIN, left: MARGIN };
      case 'bottom-right':
        return { top: h - SIZE - MARGIN, left: w - SIZE - MARGIN };
    }
  }

  private clearPositionProps() {
    this.style.top = '';
    this.style.bottom = '';
    this.style.left = '';
    this.style.right = '';
  }
}

customElements.define('yr-chat-fab', YrChatFab);
