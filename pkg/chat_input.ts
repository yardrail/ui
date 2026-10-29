export class YrChatInput extends HTMLElement {
  private staged: File[] = [];

  connectedCallback() {
    this.sendBtn?.addEventListener('click', this.handleSend);
    this.field?.addEventListener('keydown', this.handleKeydown);
    this.field?.addEventListener('input', this.handleAutoGrow);
    this.querySelector('.yr-chat-input__attach')?.addEventListener(
      'click',
      this.handleAttach,
    );
    this.fileInput?.addEventListener('change', this.handleFiles);
    this.stagingEl?.addEventListener('click', this.handleStagingClick);
  }

  disconnectedCallback() {
    this.sendBtn?.removeEventListener('click', this.handleSend);
    this.field?.removeEventListener('keydown', this.handleKeydown);
    this.field?.removeEventListener('input', this.handleAutoGrow);
    this.querySelector('.yr-chat-input__attach')?.removeEventListener(
      'click',
      this.handleAttach,
    );
    this.fileInput?.removeEventListener('change', this.handleFiles);
    this.stagingEl?.removeEventListener('click', this.handleStagingClick);
  }

  private get field(): HTMLTextAreaElement | null {
    return this.querySelector('.yr-chat-input__field');
  }

  private get sendBtn(): HTMLElement | null {
    return this.querySelector('.yr-chat-input__send');
  }

  private get fileInput(): HTMLInputElement | null {
    return this.querySelector('.yr-chat-input__file');
  }

  private get stagingEl(): HTMLElement | null {
    return this.querySelector('.yr-chat-input__staging');
  }

  private handleAutoGrow = () => {
    const el = this.field;
    if (!el) return;
    el.style.height = 'auto';
    el.style.height = `${el.scrollHeight}px`;
  };

  private handleSend = () => {
    const value = this.field?.value.trim();
    if (!value && this.staged.length === 0) return;
    this.dispatchEvent(
      new CustomEvent('yr-chat-send', {
        bubbles: true,
        detail: { text: value || '', files: [...this.staged] },
      }),
    );
    if (this.field) {
      this.field.value = '';
      this.field.style.height = 'auto';
    }
    this.staged = [];
    this.renderStaging();
  };

  private handleKeydown = (e: KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      this.handleSend();
    }
  };

  private handleAttach = () => {
    this.fileInput?.click();
  };

  private handleFiles = () => {
    const files = this.fileInput?.files;
    if (!files) return;
    for (let i = 0; i < files.length; i++) {
      this.staged.push(files[i]);
    }
    if (this.fileInput) this.fileInput.value = '';
    this.renderStaging();
  };

  private handleStagingClick = (e: Event) => {
    const btn = (e.target as HTMLElement).closest(
      '.yr-chat-input__chip-remove',
    );
    if (!btn) return;
    const idx = Number(btn.getAttribute('data-index'));
    if (!Number.isNaN(idx)) {
      this.staged.splice(idx, 1);
      this.renderStaging();
    }
  };

  private renderStaging() {
    const el = this.stagingEl;
    if (!el) return;
    if (this.staged.length === 0) {
      el.hidden = true;
      el.innerHTML = '';
      return;
    }
    el.hidden = false;
    el.innerHTML = this.staged.map((f, i) => this.chipHTML(f, i)).join('');
  }

  private chipHTML(file: File, index: number): string {
    const name =
      file.name.length > 20 ? `${file.name.slice(0, 17)}...` : file.name;
    const isImage = file.type.startsWith('image/');
    const isAudio = file.type.startsWith('audio/');
    const icon = isImage
      ? `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="9" cy="9" r="2"/><path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></svg>`
      : isAudio
        ? `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 22h.5a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"/><path d="M14 2v5h5"/><path d="M2 19a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0v-4a6 6 0 0 1 12 0v4a2 2 0 1 1-4 0v-1a2 2 0 1 1 4 0"/></svg>`
        : `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></svg>`;
    const removeIcon = `<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`;
    return `<div class="yr-chat-input__chip"><span class="yr-chat-input__chip-icon">${icon}</span><span class="yr-chat-input__chip-name">${this.escapeHTML(name)}</span><button type="button" class="yr-chat-input__chip-remove" data-index="${index}" aria-label="Remove ${this.escapeHTML(name)}">${removeIcon}</button></div>`;
  }

  private escapeHTML(s: string): string {
    const el = document.createElement('span');
    el.textContent = s;
    return el.innerHTML;
  }
}

customElements.define('yr-chat-input', YrChatInput);
