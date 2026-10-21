package ui

import "github.com/a-h/templ"

const iconAlertTriangleSVG = `<svg  class="lucide lucide-alert-triangle"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3" />  <path d="M12 9v4" />  <path d="M12 17h.01" /> </svg>`

// IconAlertTriangle renders the Lucide alert-triangle icon.
func IconAlertTriangle() templ.Component {
	return templ.Raw(iconAlertTriangleSVG)
}

const iconChevronDownSVG = `<svg class="lucide lucide-chevron-down" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" > <path d="m6 9 6 6 6-6" /> </svg>`

// IconChevronDown renders the Lucide chevron-down icon.
func IconChevronDown() templ.Component {
	return templ.Raw(iconChevronDownSVG)
}

const iconBuilding2SVG = `<svg  class="lucide lucide-building-2"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M10 12h4" />  <path d="M10 8h4" />  <path d="M14 21v-3a2 2 0 0 0-4 0v3" />  <path d="M6 10H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-2" />  <path d="M6 21V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v16" /> </svg>`

// IconBuilding2 renders the Lucide building-2 icon.
func IconBuilding2() templ.Component {
	return templ.Raw(iconBuilding2SVG)
}

const iconCheckCircleSVG = `<svg  class="lucide lucide-check-circle"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M21.801 10A10 10 0 1 1 17 3.335" />  <path d="m9 11 3 3L22 4" /> </svg>`

// IconCheckCircle renders the Lucide check-circle icon.
func IconCheckCircle() templ.Component {
	return templ.Raw(iconCheckCircleSVG)
}

const iconCircleXSVG = `<svg  class="lucide lucide-circle-x"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <circle cx="12" cy="12" r="10" />  <path d="m15 9-6 6" />  <path d="m9 9 6 6" /> </svg>`

// IconCircleX renders the Lucide circle-x icon.
func IconCircleX() templ.Component {
	return templ.Raw(iconCircleXSVG)
}

const iconClockSVG = `<svg  class="lucide lucide-clock"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <circle cx="12" cy="12" r="10" />  <path d="M12 6v6l4 2" /> </svg>`

// IconClock renders the Lucide clock icon.
func IconClock() templ.Component {
	return templ.Raw(iconClockSVG)
}

const iconCreditCardSVG = `<svg  class="lucide lucide-credit-card"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <rect width="20" height="14" x="2" y="5" rx="2" />  <line x1="2" x2="22" y1="10" y2="10" />  <path d="M6 14h2" /> </svg>`

// IconCreditCard renders the Lucide credit-card icon.
func IconCreditCard() templ.Component {
	return templ.Raw(iconCreditCardSVG)
}

const iconDatabaseSVG = `<svg  class="lucide lucide-database"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <ellipse cx="12" cy="5" rx="9" ry="3" />  <path d="M3 5V19A9 3 0 0 0 21 19V5" />  <path d="M3 12A9 3 0 0 0 21 12" /> </svg>`

// IconDatabase renders the Lucide database icon.
func IconDatabase() templ.Component {
	return templ.Raw(iconDatabaseSVG)
}

const iconFileSVG = `<svg  class="lucide lucide-file"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z" />  <path d="M14 2v5a1 1 0 0 0 1 1h5" /> </svg>`

// IconFile renders the Lucide file icon.
func IconFile() templ.Component {
	return templ.Raw(iconFileSVG)
}

const iconFileImageSVG = `<svg  class="lucide lucide-file-image"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z" />  <path d="M14 2v5a1 1 0 0 0 1 1h5" />  <circle cx="10" cy="12" r="2" />  <path d="m20 17-1.296-1.296a2.41 2.41 0 0 0-3.408 0L9 22" /> </svg>`

// IconFileImage renders the Lucide file-image icon.
func IconFileImage() templ.Component {
	return templ.Raw(iconFileImageSVG)
}

const iconFileSpreadsheetSVG = `<svg  class="lucide lucide-file-spreadsheet"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z" />  <path d="M14 2v5a1 1 0 0 0 1 1h5" />  <path d="M8 13h2" />  <path d="M14 13h2" />  <path d="M8 17h2" />  <path d="M14 17h2" /> </svg>`

// IconFileSpreadsheet renders the Lucide file-spreadsheet icon.
func IconFileSpreadsheet() templ.Component {
	return templ.Raw(iconFileSpreadsheetSVG)
}

const iconFileTextSVG = `<svg  class="lucide lucide-file-text"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z" />  <path d="M14 2v5a1 1 0 0 0 1 1h5" />  <path d="M10 9H8" />  <path d="M16 13H8" />  <path d="M16 17H8" /> </svg>`

// IconFileText renders the Lucide file-text icon.
func IconFileText() templ.Component {
	return templ.Raw(iconFileTextSVG)
}

const iconFolderSVG = `<svg  class="lucide lucide-folder"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z" /> </svg>`

// IconFolder renders the Lucide folder icon.
func IconFolder() templ.Component {
	return templ.Raw(iconFolderSVG)
}

const iconGitBranchSVG = `<svg  class="lucide lucide-git-branch"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M15 6a9 9 0 0 0-9 9V3" />  <circle cx="18" cy="6" r="3" />  <circle cx="6" cy="18" r="3" /> </svg>`

// IconGitBranch renders the Lucide git-branch icon.
func IconGitBranch() templ.Component {
	return templ.Raw(iconGitBranchSVG)
}

const iconHomeSVG = `<svg  class="lucide lucide-home"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8" />  <path d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" /> </svg>`

// IconHome renders the Lucide home icon.
func IconHome() templ.Component {
	return templ.Raw(iconHomeSVG)
}

const iconLayoutDashboardSVG = `<svg  class="lucide lucide-layout-dashboard"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <rect width="7" height="9" x="3" y="3" rx="1" />  <rect width="7" height="5" x="14" y="3" rx="1" />  <rect width="7" height="9" x="14" y="12" rx="1" />  <rect width="7" height="5" x="3" y="16" rx="1" /> </svg>`

// IconLayoutDashboard renders the Lucide layout-dashboard icon.
func IconLayoutDashboard() templ.Component {
	return templ.Raw(iconLayoutDashboardSVG)
}

const iconLayoutGridSVG = `<svg  class="lucide lucide-layout-grid"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <rect width="7" height="7" x="3" y="3" rx="1" />  <rect width="7" height="7" x="14" y="3" rx="1" />  <rect width="7" height="7" x="14" y="14" rx="1" />  <rect width="7" height="7" x="3" y="14" rx="1" /> </svg>`

// IconLayoutGrid renders the Lucide layout-grid icon.
func IconLayoutGrid() templ.Component {
	return templ.Raw(iconLayoutGridSVG)
}

const iconListSVG = `<svg  class="lucide lucide-list"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M3 5h.01" />  <path d="M3 12h.01" />  <path d="M3 19h.01" />  <path d="M8 5h13" />  <path d="M8 12h13" />  <path d="M8 19h13" /> </svg>`

// IconList renders the Lucide list icon.
func IconList() templ.Component {
	return templ.Raw(iconListSVG)
}

const iconMailSVG = `<svg  class="lucide lucide-mail"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7" />  <rect x="2" y="4" width="20" height="16" rx="2" /> </svg>`

// IconMail renders the Lucide mail icon.
func IconMail() templ.Component {
	return templ.Raw(iconMailSVG)
}

const iconMessageSquareSVG = `<svg  class="lucide lucide-message-square"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z" /> </svg>`

// IconMessageSquare renders the Lucide message-square icon.
func IconMessageSquare() templ.Component {
	return templ.Raw(iconMessageSquareSVG)
}

const iconPencilSVG = `<svg  class="lucide lucide-pencil"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z" />  <path d="m15 5 4 4" /> </svg>`

// IconPencil renders the Lucide pencil icon.
func IconPencil() templ.Component {
	return templ.Raw(iconPencilSVG)
}

const iconPlaySVG = `<svg  class="lucide lucide-play"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M5 5a2 2 0 0 1 3.008-1.728l11.997 6.998a2 2 0 0 1 .003 3.458l-12 7A2 2 0 0 1 5 19z" /> </svg>`

// IconPlay renders the Lucide play icon.
func IconPlay() templ.Component {
	return templ.Raw(iconPlaySVG)
}

const iconPlugSVG = `<svg  class="lucide lucide-plug"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M12 22v-5" />  <path d="M15 8V2" />  <path d="M17 8a1 1 0 0 1 1 1v4a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1z" />  <path d="M9 8V2" /> </svg>`

// IconPlug renders the Lucide plug icon.
func IconPlug() templ.Component {
	return templ.Raw(iconPlugSVG)
}

const iconRefreshCwSVG = `<svg  class="lucide lucide-refresh-cw"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8" />  <path d="M21 3v5h-5" />  <path d="M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16" />  <path d="M8 16H3v5" /> </svg>`

// IconRefreshCw renders the Lucide refresh-cw icon.
func IconRefreshCw() templ.Component {
	return templ.Raw(iconRefreshCwSVG)
}

const iconSearchSVG = `<svg  class="lucide lucide-search"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="m21 21-4.34-4.34" />  <circle cx="11" cy="11" r="8" /> </svg>`

// IconSearch renders the Lucide search icon.
func IconSearch() templ.Component {
	return templ.Raw(iconSearchSVG)
}

const iconSettingsSVG = `<svg  class="lucide lucide-settings"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M9.671 4.136a2.34 2.34 0 0 1 4.659 0 2.34 2.34 0 0 0 3.319 1.915 2.34 2.34 0 0 1 2.33 4.033 2.34 2.34 0 0 0 0 3.831 2.34 2.34 0 0 1-2.33 4.033 2.34 2.34 0 0 0-3.319 1.915 2.34 2.34 0 0 1-4.659 0 2.34 2.34 0 0 0-3.32-1.915 2.34 2.34 0 0 1-2.33-4.033 2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915" />  <circle cx="12" cy="12" r="3" /> </svg>`

// IconSettings renders the Lucide settings icon.
func IconSettings() templ.Component {
	return templ.Raw(iconSettingsSVG)
}

const iconShieldSVG = `<svg  class="lucide lucide-shield"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z" /> </svg>`

// IconShield renders the Lucide shield icon.
func IconShield() templ.Component {
	return templ.Raw(iconShieldSVG)
}

const iconStarSVG = `<svg  class="lucide lucide-star"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z" /> </svg>`

// IconStar renders the Lucide star icon.
func IconStar() templ.Component {
	return templ.Raw(iconStarSVG)
}

const iconTrainFrontSVG = `<svg  class="lucide lucide-train-front"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M8 3.1V7a4 4 0 0 0 8 0V3.1" />  <path d="m9 15-1-1" />  <path d="m15 15 1-1" />  <path d="M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Z" />  <path d="m8 19-2 3" />  <path d="m16 19 2 3" /> </svg>`

// IconTrainFront renders the Lucide train-front icon.
func IconTrainFront() templ.Component {
	return templ.Raw(iconTrainFrontSVG)
}

const iconUserSVG = `<svg  class="lucide lucide-user"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />  <circle cx="12" cy="7" r="4" /> </svg>`

// IconUser renders the Lucide user icon.
func IconUser() templ.Component {
	return templ.Raw(iconUserSVG)
}

const iconUserPlusSVG = `<svg  class="lucide lucide-user-plus"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />  <circle cx="9" cy="7" r="4" />  <line x1="19" x2="19" y1="8" y2="14" />  <line x1="22" x2="16" y1="11" y2="11" /> </svg>`

// IconUserPlus renders the Lucide user-plus icon.
func IconUserPlus() templ.Component {
	return templ.Raw(iconUserPlusSVG)
}

const iconUsersSVG = `<svg  class="lucide lucide-users"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />  <path d="M16 3.128a4 4 0 0 1 0 7.744" />  <path d="M22 21v-2a4 4 0 0 0-3-3.87" />  <circle cx="9" cy="7" r="4" /> </svg>`

// IconUsers renders the Lucide users icon.
func IconUsers() templ.Component {
	return templ.Raw(iconUsersSVG)
}

const iconXSVG = `<svg  class="lucide lucide-x"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M18 6 6 18" />  <path d="m6 6 12 12" /> </svg>`

// IconX renders the Lucide x icon.
func IconX() templ.Component {
	return templ.Raw(iconXSVG)
}

const iconPanelLeftCloseSVG = `<svg class="lucide lucide-panel-left-close yr-sidebar-toggle-close" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m16 15-3-3 3-3"/></svg>`

// IconPanelLeftClose renders the sidebar collapse toggle icon.
func IconPanelLeftClose() templ.Component {
	return templ.Raw(iconPanelLeftCloseSVG)
}

const iconPanelLeftOpenSVG = `<svg class="lucide lucide-panel-left-open yr-sidebar-toggle-open" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m14 9 3 3-3 3"/></svg>`

// IconPanelLeftOpen renders the sidebar expand toggle icon.
func IconPanelLeftOpen() templ.Component {
	return templ.Raw(iconPanelLeftOpenSVG)
}

const iconZapSVG = `<svg  class="lucide lucide-zap"  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round" >  <path d="M15.914 4a1.5 1.5 0 00-2.474-1.561l-9 9A1.5 1.5 0 005.5 14h4.002a.5.5 0 01.471.666L8.086 20a1.5 1.5 0 002.475 1.56l9-9A1.5 1.5 0 0018.5 10h-3.997a.5.5 0 01-.472-.667z" /> </svg>`

// IconZap renders the Lucide zap icon.
func IconZap() templ.Component {
	return templ.Raw(iconZapSVG)
}
