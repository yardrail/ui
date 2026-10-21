import { YrDropdownMenu } from './dropdown_menu';

// <yr-account-menu> — a higher-level wrapper around yr-dropdown-menu that
// auto-builds the trigger (avatar + chevron) and panel header (avatar + name
// + email) from attributes. The chevron SVG is server-rendered in a
// [slot="chevron"] child by the templ component.
export class YrAccountMenu extends YrDropdownMenu {
  connectedCallback() {
    this.restructure();
    super.connectedCallback();
  }

  private restructure() {
    if (this.querySelector(':scope > [slot="trigger"]')) return;

    const name = this.getAttribute('name') ?? '';
    const email = this.getAttribute('email') ?? '';
    const avatarSlot = this.querySelector('[slot="avatar"]');
    const chevronSlot = this.querySelector('[slot="chevron"]');
    const triggerExtra = this.querySelector('[slot="trigger-extra"]');
    const items = Array.from(this.childNodes).filter(
      (node) => node !== avatarSlot && node !== chevronSlot && node !== triggerExtra,
    );

    const trigger = document.createElement('button');
    trigger.setAttribute('slot', 'trigger');
    trigger.className = 'yr-account-menu-trigger';
    if (triggerExtra) trigger.append(triggerExtra.cloneNode(true));
    const triggerAvatar = document.createElement('yr-avatar');
    if (avatarSlot) triggerAvatar.append(avatarSlot.cloneNode(true));
    trigger.append(triggerAvatar);
    if (chevronSlot) trigger.append(chevronSlot);

    const header = document.createElement('div');
    header.className = 'yr-dropdown-menu-header';
    const headerAvatar = document.createElement('yr-avatar');
    headerAvatar.setAttribute('size', 'lg');
    if (avatarSlot) headerAvatar.append(avatarSlot.cloneNode(true));
    header.append(headerAvatar);

    if (name || email) {
      const headerText = document.createElement('div');
      if (name) {
        const nameEl = document.createElement('div');
        nameEl.className = 'yr-dropdown-menu-header-name';
        nameEl.textContent = name;
        headerText.append(nameEl);
      }
      if (email) {
        const emailEl = document.createElement('div');
        emailEl.className = 'yr-dropdown-menu-header-email';
        emailEl.textContent = email;
        headerText.append(emailEl);
      }
      header.append(headerText);
    }

    const panel = document.createElement('div');
    panel.setAttribute('slot', 'panel');
    panel.append(header, ...items);

    this.replaceChildren(trigger, panel);
  }
}

customElements.define('yr-account-menu', YrAccountMenu);
