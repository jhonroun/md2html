package main

const favicon = `data:image/png;base64,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`

const CSS = `

@font-face {
	font-family: octicons-anchor;
	src: url(data:font/woff;charset=utf-8;base64,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) format('woff');
}

body {
	background-color: #0d1117;
}

.readme {
	width: 790px;
	display: table;
	margin: 20px auto 20px auto;
	border-radius: 3px;
	border: 1px solid #2a2a2a;
	padding: 30px;
}

.markdown-body {
	-ms-text-size-adjust: 100%;
	-webkit-text-size-adjust: 100%;
	color: #333;
	overflow: hidden;
	font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif;
	font-size: 16px;
	line-height: 1.6;
	word-wrap: break-word;
}

@media (prefers-color-scheme: dark) {
	.markdown-body,
	[data-theme="dark"] {
	  /*dark*/
	  color-scheme: dark;
	  --color-prettylights-syntax-comment: #8b949e;
	  --color-prettylights-syntax-constant: #79c0ff;
	  --color-prettylights-syntax-entity: #d2a8ff;
	  --color-prettylights-syntax-storage-modifier-import: #c9d1d9;
	  --color-prettylights-syntax-entity-tag: #7ee787;
	  --color-prettylights-syntax-keyword: #ff7b72;
	  --color-prettylights-syntax-string: #a5d6ff;
	  --color-prettylights-syntax-variable: #ffa657;
	  --color-prettylights-syntax-brackethighlighter-unmatched: #f85149;
	  --color-prettylights-syntax-invalid-illegal-text: #f0f6fc;
	  --color-prettylights-syntax-invalid-illegal-bg: #8e1519;
	  --color-prettylights-syntax-carriage-return-text: #f0f6fc;
	  --color-prettylights-syntax-carriage-return-bg: #b62324;
	  --color-prettylights-syntax-string-regexp: #7ee787;
	  --color-prettylights-syntax-markup-list: #f2cc60;
	  --color-prettylights-syntax-markup-heading: #1f6feb;
	  --color-prettylights-syntax-markup-italic: #c9d1d9;
	  --color-prettylights-syntax-markup-bold: #c9d1d9;
	  --color-prettylights-syntax-markup-deleted-text: #ffdcd7;
	  --color-prettylights-syntax-markup-deleted-bg: #67060c;
	  --color-prettylights-syntax-markup-inserted-text: #aff5b4;
	  --color-prettylights-syntax-markup-inserted-bg: #033a16;
	  --color-prettylights-syntax-markup-changed-text: #ffdfb6;
	  --color-prettylights-syntax-markup-changed-bg: #5a1e02;
	  --color-prettylights-syntax-markup-ignored-text: #c9d1d9;
	  --color-prettylights-syntax-markup-ignored-bg: #1158c7;
	  --color-prettylights-syntax-meta-diff-range: #d2a8ff;
	  --color-prettylights-syntax-brackethighlighter-angle: #8b949e;
	  --color-prettylights-syntax-sublimelinter-gutter-mark: #484f58;
	  --color-prettylights-syntax-constant-other-reference-link: #a5d6ff;
	  --color-fg-default: #e6edf3;
	  --color-fg-muted: #848d97;
	  --color-fg-subtle: #6e7681;
	  --color-canvas-default: #0d1117;
	  --color-canvas-subtle: #161b22;
	  --color-border-default: #30363d;
	  --color-border-muted: #21262d;
	  --color-neutral-muted: rgba(110,118,129,0.4);
	  --color-accent-fg: #2f81f7;
	  --color-accent-emphasis: #1f6feb;
	  --color-success-fg: #3fb950;
	  --color-success-emphasis: #238636;
	  --color-attention-fg: #d29922;
	  --color-attention-emphasis: #9e6a03;
	  --color-attention-subtle: rgba(187,128,9,0.15);
	  --color-danger-fg: #f85149;
	  --color-danger-emphasis: #da3633;
	  --color-done-fg: #a371f7;
	  --color-done-emphasis: #8957e5;
	}
  }
  
  @media (prefers-color-scheme: light) {
	.markdown-body,
	[data-theme="light"] {
	  /*light*/
	  color-scheme: light;
	  --color-prettylights-syntax-comment: #57606a;
	  --color-prettylights-syntax-constant: #0550ae;
	  --color-prettylights-syntax-entity: #6639ba;
	  --color-prettylights-syntax-storage-modifier-import: #24292f;
	  --color-prettylights-syntax-entity-tag: #116329;
	  --color-prettylights-syntax-keyword: #cf222e;
	  --color-prettylights-syntax-string: #0a3069;
	  --color-prettylights-syntax-variable: #953800;
	  --color-prettylights-syntax-brackethighlighter-unmatched: #82071e;
	  --color-prettylights-syntax-invalid-illegal-text: #f6f8fa;
	  --color-prettylights-syntax-invalid-illegal-bg: #82071e;
	  --color-prettylights-syntax-carriage-return-text: #f6f8fa;
	  --color-prettylights-syntax-carriage-return-bg: #cf222e;
	  --color-prettylights-syntax-string-regexp: #116329;
	  --color-prettylights-syntax-markup-list: #3b2300;
	  --color-prettylights-syntax-markup-heading: #0550ae;
	  --color-prettylights-syntax-markup-italic: #24292f;
	  --color-prettylights-syntax-markup-bold: #24292f;
	  --color-prettylights-syntax-markup-deleted-text: #82071e;
	  --color-prettylights-syntax-markup-deleted-bg: #ffebe9;
	  --color-prettylights-syntax-markup-inserted-text: #116329;
	  --color-prettylights-syntax-markup-inserted-bg: #dafbe1;
	  --color-prettylights-syntax-markup-changed-text: #953800;
	  --color-prettylights-syntax-markup-changed-bg: #ffd8b5;
	  --color-prettylights-syntax-markup-ignored-text: #eaeef2;
	  --color-prettylights-syntax-markup-ignored-bg: #0550ae;
	  --color-prettylights-syntax-meta-diff-range: #8250df;
	  --color-prettylights-syntax-brackethighlighter-angle: #57606a;
	  --color-prettylights-syntax-sublimelinter-gutter-mark: #8c959f;
	  --color-prettylights-syntax-constant-other-reference-link: #0a3069;
	  --color-fg-default: #1F2328;
	  --color-fg-muted: #656d76;
	  --color-fg-subtle: #6e7781;
	  --color-canvas-default: #ffffff;
	  --color-canvas-subtle: #f6f8fa;
	  --color-border-default: #d0d7de;
	  --color-border-muted: hsla(210,18%,87%,1);
	  --color-neutral-muted: rgba(175,184,193,0.2);
	  --color-accent-fg: #0969da;
	  --color-accent-emphasis: #0969da;
	  --color-success-fg: #1a7f37;
	  --color-success-emphasis: #1f883d;
	  --color-attention-fg: #9a6700;
	  --color-attention-emphasis: #9a6700;
	  --color-attention-subtle: #fff8c5;
	  --color-danger-fg: #d1242f;
	  --color-danger-emphasis: #cf222e;
	  --color-done-fg: #8250df;
	  --color-done-emphasis: #8250df;
	}
  }
  
  .markdown-body {
	-ms-text-size-adjust: 100%;
	-webkit-text-size-adjust: 100%;
	margin: 0;
	color: var(--color-fg-default);
	background-color: var(--color-canvas-default);
	font-family: -apple-system,BlinkMacSystemFont,"Segoe UI","Noto Sans",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji";
	font-size: 16px;
	line-height: 1.5;
	word-wrap: break-word;
  }
  
  .markdown-body .octicon {
	display: inline-block;
	fill: currentColor;
	vertical-align: text-bottom;
  }
  
  .markdown-body h1:hover .anchor .octicon-link:before,
  .markdown-body h2:hover .anchor .octicon-link:before,
  .markdown-body h3:hover .anchor .octicon-link:before,
  .markdown-body h4:hover .anchor .octicon-link:before,
  .markdown-body h5:hover .anchor .octicon-link:before,
  .markdown-body h6:hover .anchor .octicon-link:before {
	width: 16px;
	height: 16px;
	content: ' ';
	display: inline-block;
	background-color: currentColor;
	-webkit-mask-image: url("data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' version='1.1' aria-hidden='true'><path fill-rule='evenodd' d='M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z'></path></svg>");
	mask-image: url("data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' version='1.1' aria-hidden='true'><path fill-rule='evenodd' d='M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z'></path></svg>");
  }
  
  .markdown-body details,
  .markdown-body figcaption,
  .markdown-body figure {
	display: block;
  }
  
  .markdown-body summary {
	display: list-item;
  }
  
  .markdown-body [hidden] {
	display: none !important;
  }
  
  .markdown-body a {
	background-color: transparent;
	color: var(--color-accent-fg);
	text-decoration: none;
  }
  
  .markdown-body abbr[title] {
	border-bottom: none;
	-webkit-text-decoration: underline dotted;
	text-decoration: underline dotted;
  }
  
  .markdown-body b,
  .markdown-body strong {
	font-weight: var(--base-text-weight-semibold, 600);
  }
  
  .markdown-body dfn {
	font-style: italic;
  }
  
  .markdown-body h1 {
	margin: .67em 0;
	font-weight: var(--base-text-weight-semibold, 600);
	padding-bottom: .3em;
	font-size: 2em;
	border-bottom: 1px solid var(--color-border-muted);
  }
  
  .markdown-body mark {
	background-color: var(--color-attention-subtle);
	color: var(--color-fg-default);
  }
  
  .markdown-body small {
	font-size: 90%;
  }
  
  .markdown-body sub,
  .markdown-body sup {
	font-size: 75%;
	line-height: 0;
	position: relative;
	vertical-align: baseline;
  }
  
  .markdown-body sub {
	bottom: -0.25em;
  }
  
  .markdown-body sup {
	top: -0.5em;
  }
  
  .markdown-body img {
	border-style: none;
	max-width: 100%;
	box-sizing: content-box;
	background-color: var(--color-canvas-default);
  }
  
  .markdown-body code,
  .markdown-body kbd,
  .markdown-body pre,
  .markdown-body samp {
	font-family: monospace;
	font-size: 1em;
  }
  
  .markdown-body figure {
	margin: 1em 40px;
  }
  
  .markdown-body hr {
	box-sizing: content-box;
	overflow: hidden;
	background: transparent;
	border-bottom: 1px solid var(--color-border-muted);
	height: .25em;
	padding: 0;
	margin: 24px 0;
	background-color: var(--color-border-default);
	border: 0;
  }
  
  .markdown-body input {
	font: inherit;
	margin: 0;
	overflow: visible;
	font-family: inherit;
	font-size: inherit;
	line-height: inherit;
  }
  
  .markdown-body [type=button],
  .markdown-body [type=reset],
  .markdown-body [type=submit] {
	-webkit-appearance: button;
	appearance: button;
  }
  
  .markdown-body [type=checkbox],
  .markdown-body [type=radio] {
	box-sizing: border-box;
	padding: 0;
  }
  
  .markdown-body [type=number]::-webkit-inner-spin-button,
  .markdown-body [type=number]::-webkit-outer-spin-button {
	height: auto;
  }
  
  .markdown-body [type=search]::-webkit-search-cancel-button,
  .markdown-body [type=search]::-webkit-search-decoration {
	-webkit-appearance: none;
	appearance: none;
  }
  
  .markdown-body ::-webkit-input-placeholder {
	color: inherit;
	opacity: .54;
  }
  
  .markdown-body ::-webkit-file-upload-button {
	-webkit-appearance: button;
	appearance: button;
	font: inherit;
  }
  
  .markdown-body a:hover {
	text-decoration: underline;
  }
  
  .markdown-body ::placeholder {
	color: var(--color-fg-subtle);
	opacity: 1;
  }
  
  .markdown-body hr::before {
	display: table;
	content: "";
  }
  
  .markdown-body hr::after {
	display: table;
	clear: both;
	content: "";
  }
  
  .markdown-body table {
	border-spacing: 0;
	border-collapse: collapse;
	display: block;
	width: max-content;
	max-width: 100%;
	overflow: auto;
  }
  
  .markdown-body td,
  .markdown-body th {
	padding: 0;
  }
  
  .markdown-body details summary {
	cursor: pointer;
  }
  
  .markdown-body details:not([open])>*:not(summary) {
	display: none !important;
  }
  
  .markdown-body a:focus,
  .markdown-body [role=button]:focus,
  .markdown-body input[type=radio]:focus,
  .markdown-body input[type=checkbox]:focus {
	outline: 2px solid var(--color-accent-fg);
	outline-offset: -2px;
	box-shadow: none;
  }
  
  .markdown-body a:focus:not(:focus-visible),
  .markdown-body [role=button]:focus:not(:focus-visible),
  .markdown-body input[type=radio]:focus:not(:focus-visible),
  .markdown-body input[type=checkbox]:focus:not(:focus-visible) {
	outline: solid 1px transparent;
  }
  
  .markdown-body a:focus-visible,
  .markdown-body [role=button]:focus-visible,
  .markdown-body input[type=radio]:focus-visible,
  .markdown-body input[type=checkbox]:focus-visible {
	outline: 2px solid var(--color-accent-fg);
	outline-offset: -2px;
	box-shadow: none;
  }
  
  .markdown-body a:not([class]):focus,
  .markdown-body a:not([class]):focus-visible,
  .markdown-body input[type=radio]:focus,
  .markdown-body input[type=radio]:focus-visible,
  .markdown-body input[type=checkbox]:focus,
  .markdown-body input[type=checkbox]:focus-visible {
	outline-offset: 0;
  }
  
  .markdown-body kbd {
	display: inline-block;
	padding: 3px 5px;
	font: 11px ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
	line-height: 10px;
	color: var(--color-fg-default);
	vertical-align: middle;
	background-color: var(--color-canvas-subtle);
	border: solid 1px var(--color-neutral-muted);
	border-bottom-color: var(--color-neutral-muted);
	border-radius: 6px;
	box-shadow: inset 0 -1px 0 var(--color-neutral-muted);
  }
  
  .markdown-body h1,
  .markdown-body h2,
  .markdown-body h3,
  .markdown-body h4,
  .markdown-body h5,
  .markdown-body h6 {
	margin-top: 24px;
	margin-bottom: 16px;
	font-weight: var(--base-text-weight-semibold, 600);
	line-height: 1.25;
  }
  
  .markdown-body h2 {
	font-weight: var(--base-text-weight-semibold, 600);
	padding-bottom: .3em;
	font-size: 1.5em;
	border-bottom: 1px solid var(--color-border-muted);
  }
  
  .markdown-body h3 {
	font-weight: var(--base-text-weight-semibold, 600);
	font-size: 1.25em;
  }
  
  .markdown-body h4 {
	font-weight: var(--base-text-weight-semibold, 600);
	font-size: 1em;
  }
  
  .markdown-body h5 {
	font-weight: var(--base-text-weight-semibold, 600);
	font-size: .875em;
  }
  
  .markdown-body h6 {
	font-weight: var(--base-text-weight-semibold, 600);
	font-size: .85em;
	color: var(--color-fg-muted);
  }
  
  .markdown-body p {
	margin-top: 0;
	margin-bottom: 10px;
  }
  
  .markdown-body blockquote {
	margin: 0;
	padding: 0 1em;
	color: var(--color-fg-muted);
	border-left: .25em solid var(--color-border-default);
  }
  
  .markdown-body ul,
  .markdown-body ol {
	margin-top: 0;
	margin-bottom: 0;
	padding-left: 2em;
  }
  
  .markdown-body ol ol,
  .markdown-body ul ol {
	list-style-type: lower-roman;
  }
  
  .markdown-body ul ul ol,
  .markdown-body ul ol ol,
  .markdown-body ol ul ol,
  .markdown-body ol ol ol {
	list-style-type: lower-alpha;
  }
  
  .markdown-body dd {
	margin-left: 0;
  }
  
  .markdown-body tt,
  .markdown-body code,
  .markdown-body samp {
	font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
	font-size: 12px;
  }
  
  .markdown-body pre {
	margin-top: 0;
	margin-bottom: 0;
	font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
	font-size: 12px;
	word-wrap: normal;
  }
  
  .markdown-body .octicon {
	display: inline-block;
	overflow: visible !important;
	vertical-align: text-bottom;
	fill: currentColor;
  }
  
  .markdown-body input::-webkit-outer-spin-button,
  .markdown-body input::-webkit-inner-spin-button {
	margin: 0;
	-webkit-appearance: none;
	appearance: none;
  }
  
  .markdown-body .mr-2 {
	margin-right: var(--base-size-8, 8px) !important;
  }
  
  .markdown-body::before {
	display: table;
	content: "";
  }
  
  .markdown-body::after {
	display: table;
	clear: both;
	content: "";
  }
  
  .markdown-body>*:first-child {
	margin-top: 0 !important;
  }
  
  .markdown-body>*:last-child {
	margin-bottom: 0 !important;
  }
  
  .markdown-body a:not([href]) {
	color: inherit;
	text-decoration: none;
  }
  
  .markdown-body .absent {
	color: var(--color-danger-fg);
  }
  
  .markdown-body .anchor {
	float: left;
	padding-right: 4px;
	margin-left: -20px;
	line-height: 1;
  }
  
  .markdown-body .anchor:focus {
	outline: none;
  }
  
  .markdown-body p,
  .markdown-body blockquote,
  .markdown-body ul,
  .markdown-body ol,
  .markdown-body dl,
  .markdown-body table,
  .markdown-body pre,
  .markdown-body details {
	margin-top: 0;
	margin-bottom: 16px;
  }
  
  .markdown-body blockquote>:first-child {
	margin-top: 0;
  }
  
  .markdown-body blockquote>:last-child {
	margin-bottom: 0;
  }
  
  .markdown-body h1 .octicon-link,
  .markdown-body h2 .octicon-link,
  .markdown-body h3 .octicon-link,
  .markdown-body h4 .octicon-link,
  .markdown-body h5 .octicon-link,
  .markdown-body h6 .octicon-link {
	color: var(--color-fg-default);
	vertical-align: middle;
	visibility: hidden;
  }
  
  .markdown-body h1:hover .anchor,
  .markdown-body h2:hover .anchor,
  .markdown-body h3:hover .anchor,
  .markdown-body h4:hover .anchor,
  .markdown-body h5:hover .anchor,
  .markdown-body h6:hover .anchor {
	text-decoration: none;
  }
  
  .markdown-body h1:hover .anchor .octicon-link,
  .markdown-body h2:hover .anchor .octicon-link,
  .markdown-body h3:hover .anchor .octicon-link,
  .markdown-body h4:hover .anchor .octicon-link,
  .markdown-body h5:hover .anchor .octicon-link,
  .markdown-body h6:hover .anchor .octicon-link {
	visibility: visible;
  }
  
  .markdown-body h1 tt,
  .markdown-body h1 code,
  .markdown-body h2 tt,
  .markdown-body h2 code,
  .markdown-body h3 tt,
  .markdown-body h3 code,
  .markdown-body h4 tt,
  .markdown-body h4 code,
  .markdown-body h5 tt,
  .markdown-body h5 code,
  .markdown-body h6 tt,
  .markdown-body h6 code {
	padding: 0 .2em;
	font-size: inherit;
  }
  
  .markdown-body summary h1,
  .markdown-body summary h2,
  .markdown-body summary h3,
  .markdown-body summary h4,
  .markdown-body summary h5,
  .markdown-body summary h6 {
	display: inline-block;
  }
  
  .markdown-body summary h1 .anchor,
  .markdown-body summary h2 .anchor,
  .markdown-body summary h3 .anchor,
  .markdown-body summary h4 .anchor,
  .markdown-body summary h5 .anchor,
  .markdown-body summary h6 .anchor {
	margin-left: -40px;
  }
  
  .markdown-body summary h1,
  .markdown-body summary h2 {
	padding-bottom: 0;
	border-bottom: 0;
  }
  
  .markdown-body ul.no-list,
  .markdown-body ol.no-list {
	padding: 0;
	list-style-type: none;
  }
  
  .markdown-body ol[type="a s"] {
	list-style-type: lower-alpha;
  }
  
  .markdown-body ol[type="A s"] {
	list-style-type: upper-alpha;
  }
  
  .markdown-body ol[type="i s"] {
	list-style-type: lower-roman;
  }
  
  .markdown-body ol[type="I s"] {
	list-style-type: upper-roman;
  }
  
  .markdown-body ol[type="1"] {
	list-style-type: decimal;
  }
  
  .markdown-body div>ol:not([type]) {
	list-style-type: decimal;
  }
  
  .markdown-body ul ul,
  .markdown-body ul ol,
  .markdown-body ol ol,
  .markdown-body ol ul {
	margin-top: 0;
	margin-bottom: 0;
  }
  
  .markdown-body li>p {
	margin-top: 16px;
  }
  
  .markdown-body li+li {
	margin-top: .25em;
  }
  
  .markdown-body dl {
	padding: 0;
  }
  
  .markdown-body dl dt {
	padding: 0;
	margin-top: 16px;
	font-size: 1em;
	font-style: italic;
	font-weight: var(--base-text-weight-semibold, 600);
  }
  
  .markdown-body dl dd {
	padding: 0 16px;
	margin-bottom: 16px;
  }
  
  .markdown-body table th {
	font-weight: var(--base-text-weight-semibold, 600);
  }
  
  .markdown-body table th,
  .markdown-body table td {
	padding: 6px 13px;
	border: 1px solid var(--color-border-default);
  }
  
  .markdown-body table td>:last-child {
	margin-bottom: 0;
  }
  
  .markdown-body table tr {
	background-color: var(--color-canvas-default);
	border-top: 1px solid var(--color-border-muted);
  }
  
  .markdown-body table tr:nth-child(2n) {
	background-color: var(--color-canvas-subtle);
  }
  
  .markdown-body table img {
	background-color: transparent;
  }
  
  .markdown-body img[align=right] {
	padding-left: 20px;
  }
  
  .markdown-body img[align=left] {
	padding-right: 20px;
  }
  
  .markdown-body .emoji {
	max-width: none;
	vertical-align: text-top;
	background-color: transparent;
  }
  
  .markdown-body span.frame {
	display: block;
	overflow: hidden;
  }
  
  .markdown-body span.frame>span {
	display: block;
	float: left;
	width: auto;
	padding: 7px;
	margin: 13px 0 0;
	overflow: hidden;
	border: 1px solid var(--color-border-default);
  }
  
  .markdown-body span.frame span img {
	display: block;
	float: left;
  }
  
  .markdown-body span.frame span span {
	display: block;
	padding: 5px 0 0;
	clear: both;
	color: var(--color-fg-default);
  }
  
  .markdown-body span.align-center {
	display: block;
	overflow: hidden;
	clear: both;
  }
  
  .markdown-body span.align-center>span {
	display: block;
	margin: 13px auto 0;
	overflow: hidden;
	text-align: center;
  }
  
  .markdown-body span.align-center span img {
	margin: 0 auto;
	text-align: center;
  }
  
  .markdown-body span.align-right {
	display: block;
	overflow: hidden;
	clear: both;
  }
  
  .markdown-body span.align-right>span {
	display: block;
	margin: 13px 0 0;
	overflow: hidden;
	text-align: right;
  }
  
  .markdown-body span.align-right span img {
	margin: 0;
	text-align: right;
  }
  
  .markdown-body span.float-left {
	display: block;
	float: left;
	margin-right: 13px;
	overflow: hidden;
  }
  
  .markdown-body span.float-left span {
	margin: 13px 0 0;
  }
  
  .markdown-body span.float-right {
	display: block;
	float: right;
	margin-left: 13px;
	overflow: hidden;
  }
  
  .markdown-body span.float-right>span {
	display: block;
	margin: 13px auto 0;
	overflow: hidden;
	text-align: right;
  }
  
  .markdown-body code,
  .markdown-body tt {
	padding: .2em .4em;
	margin: 0;
	font-size: 85%;
	white-space: break-spaces;
	background-color: var(--color-neutral-muted);
	border-radius: 6px;
  }
  
  .markdown-body code br,
  .markdown-body tt br {
	display: none;
  }
  
  .markdown-body del code {
	text-decoration: inherit;
  }
  
  .markdown-body samp {
	font-size: 85%;
  }
  
  .markdown-body pre code {
	font-size: 100%;
  }
  
  .markdown-body pre>code {
	padding: 0;
	margin: 0;
	word-break: normal;
	white-space: pre;
	background: transparent;
	border: 0;
  }
  
  .markdown-body .highlight {
	margin-bottom: 16px;
  }
  
  .markdown-body .highlight pre {
	margin-bottom: 0;
	word-break: normal;
  }
  
  .markdown-body .highlight pre,
  .markdown-body pre {
	padding: 16px;
	overflow: auto;
	font-size: 85%;
	line-height: 1.45;
	color: var(--color-fg-default);
	background-color: var(--color-canvas-subtle);
	border-radius: 6px;
  }
  
  .markdown-body pre code,
  .markdown-body pre tt {
	display: inline;
	max-width: auto;
	padding: 0;
	margin: 0;
	overflow: visible;
	line-height: inherit;
	word-wrap: normal;
	background-color: transparent;
	border: 0;
  }
  
  .markdown-body .csv-data td,
  .markdown-body .csv-data th {
	padding: 5px;
	overflow: hidden;
	font-size: 12px;
	line-height: 1;
	text-align: left;
	white-space: nowrap;
  }
  
  .markdown-body .csv-data .blob-num {
	padding: 10px 8px 9px;
	text-align: right;
	background: var(--color-canvas-default);
	border: 0;
  }
  
  .markdown-body .csv-data tr {
	border-top: 0;
  }
  
  .markdown-body .csv-data th {
	font-weight: var(--base-text-weight-semibold, 600);
	background: var(--color-canvas-subtle);
	border-top: 0;
  }
  
  .markdown-body [data-footnote-ref]::before {
	content: "[";
  }
  
  .markdown-body [data-footnote-ref]::after {
	content: "]";
  }
  
  .markdown-body .footnotes {
	font-size: 12px;
	color: var(--color-fg-muted);
	border-top: 1px solid var(--color-border-default);
  }
  
  .markdown-body .footnotes ol {
	padding-left: 16px;
  }
  
  .markdown-body .footnotes ol ul {
	display: inline-block;
	padding-left: 16px;
	margin-top: 16px;
  }
  
  .markdown-body .footnotes li {
	position: relative;
  }
  
  .markdown-body .footnotes li:target::before {
	position: absolute;
	top: -8px;
	right: -8px;
	bottom: -8px;
	left: -24px;
	pointer-events: none;
	content: "";
	border: 2px solid var(--color-accent-emphasis);
	border-radius: 6px;
  }
  
  .markdown-body .footnotes li:target {
	color: var(--color-fg-default);
  }
  
  .markdown-body .footnotes .data-footnote-backref g-emoji {
	font-family: monospace;
  }
  
  .markdown-body .pl-c {
	color: var(--color-prettylights-syntax-comment);
  }
  
  .markdown-body .pl-c1,
  .markdown-body .pl-s .pl-v {
	color: var(--color-prettylights-syntax-constant);
  }
  
  .markdown-body .pl-e,
  .markdown-body .pl-en {
	color: var(--color-prettylights-syntax-entity);
  }
  
  .markdown-body .pl-smi,
  .markdown-body .pl-s .pl-s1 {
	color: var(--color-prettylights-syntax-storage-modifier-import);
  }
  
  .markdown-body .pl-ent {
	color: var(--color-prettylights-syntax-entity-tag);
  }
  
  .markdown-body .pl-k {
	color: var(--color-prettylights-syntax-keyword);
  }
  
  .markdown-body .pl-s,
  .markdown-body .pl-pds,
  .markdown-body .pl-s .pl-pse .pl-s1,
  .markdown-body .pl-sr,
  .markdown-body .pl-sr .pl-cce,
  .markdown-body .pl-sr .pl-sre,
  .markdown-body .pl-sr .pl-sra {
	color: var(--color-prettylights-syntax-string);
  }
  
  .markdown-body .pl-v,
  .markdown-body .pl-smw {
	color: var(--color-prettylights-syntax-variable);
  }
  
  .markdown-body .pl-bu {
	color: var(--color-prettylights-syntax-brackethighlighter-unmatched);
  }
  
  .markdown-body .pl-ii {
	color: var(--color-prettylights-syntax-invalid-illegal-text);
	background-color: var(--color-prettylights-syntax-invalid-illegal-bg);
  }
  
  .markdown-body .pl-c2 {
	color: var(--color-prettylights-syntax-carriage-return-text);
	background-color: var(--color-prettylights-syntax-carriage-return-bg);
  }
  
  .markdown-body .pl-sr .pl-cce {
	font-weight: bold;
	color: var(--color-prettylights-syntax-string-regexp);
  }
  
  .markdown-body .pl-ml {
	color: var(--color-prettylights-syntax-markup-list);
  }
  
  .markdown-body .pl-mh,
  .markdown-body .pl-mh .pl-en,
  .markdown-body .pl-ms {
	font-weight: bold;
	color: var(--color-prettylights-syntax-markup-heading);
  }
  
  .markdown-body .pl-mi {
	font-style: italic;
	color: var(--color-prettylights-syntax-markup-italic);
  }
  
  .markdown-body .pl-mb {
	font-weight: bold;
	color: var(--color-prettylights-syntax-markup-bold);
  }
  
  .markdown-body .pl-md {
	color: var(--color-prettylights-syntax-markup-deleted-text);
	background-color: var(--color-prettylights-syntax-markup-deleted-bg);
  }
  
  .markdown-body .pl-mi1 {
	color: var(--color-prettylights-syntax-markup-inserted-text);
	background-color: var(--color-prettylights-syntax-markup-inserted-bg);
  }
  
  .markdown-body .pl-mc {
	color: var(--color-prettylights-syntax-markup-changed-text);
	background-color: var(--color-prettylights-syntax-markup-changed-bg);
  }
  
  .markdown-body .pl-mi2 {
	color: var(--color-prettylights-syntax-markup-ignored-text);
	background-color: var(--color-prettylights-syntax-markup-ignored-bg);
  }
  
  .markdown-body .pl-mdr {
	font-weight: bold;
	color: var(--color-prettylights-syntax-meta-diff-range);
  }
  
  .markdown-body .pl-ba {
	color: var(--color-prettylights-syntax-brackethighlighter-angle);
  }
  
  .markdown-body .pl-sg {
	color: var(--color-prettylights-syntax-sublimelinter-gutter-mark);
  }
  
  .markdown-body .pl-corl {
	text-decoration: underline;
	color: var(--color-prettylights-syntax-constant-other-reference-link);
  }
  
  .markdown-body g-emoji {
	display: inline-block;
	min-width: 1ch;
	font-family: "Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";
	font-size: 1em;
	font-style: normal !important;
	font-weight: var(--base-text-weight-normal, 400);
	line-height: 1;
	vertical-align: -0.075em;
  }
  
  .markdown-body g-emoji img {
	width: 1em;
	height: 1em;
  }
  
  .markdown-body .task-list-item {
	list-style-type: none;
  }
  
  .markdown-body .task-list-item label {
	font-weight: var(--base-text-weight-normal, 400);
  }
  
  .markdown-body .task-list-item.enabled label {
	cursor: pointer;
  }
  
  .markdown-body .task-list-item+.task-list-item {
	margin-top: 4px;
  }
  
  .markdown-body .task-list-item .handle {
	display: none;
  }
  
  .markdown-body .task-list-item-checkbox {
	margin: 0 .2em .25em -1.4em;
	vertical-align: middle;
  }
  
  .markdown-body .contains-task-list:dir(rtl) .task-list-item-checkbox {
	margin: 0 -1.6em .25em .2em;
  }
  
  .markdown-body .contains-task-list {
	position: relative;
  }
  
  .markdown-body .contains-task-list:hover .task-list-item-convert-container,
  .markdown-body .contains-task-list:focus-within .task-list-item-convert-container {
	display: block;
	width: auto;
	height: 24px;
	overflow: visible;
	clip: auto;
  }
  
  .markdown-body ::-webkit-calendar-picker-indicator {
	filter: invert(50%);
  }
  
  .markdown-body .markdown-alert {
	padding: var(--base-size-8) var(--base-size-16);
	margin-bottom: 16px;
	color: inherit;
	border-left: .25em solid var(--color-border-default);
  }
  
  .markdown-body .markdown-alert>:first-child {
	margin-top: 0;
  }
  
  .markdown-body .markdown-alert>:last-child {
	margin-bottom: 0;
  }
  
  .markdown-body .markdown-alert .markdown-alert-title {
	display: flex;
	font-weight: var(--base-text-weight-medium, 500);
	align-items: center;
	line-height: 1;
  }
  
  .markdown-body .markdown-alert.markdown-alert-note {
	border-left-color: var(--color-accent-emphasis);
  }
  
  .markdown-body .markdown-alert.markdown-alert-note .markdown-alert-title {
	color: var(--color-accent-fg);
  }
  
  .markdown-body .markdown-alert.markdown-alert-important {
	border-left-color: var(--color-done-emphasis);
  }
  
  .markdown-body .markdown-alert.markdown-alert-important .markdown-alert-title {
	color: var(--color-done-fg);
  }
  
  .markdown-body .markdown-alert.markdown-alert-warning {
	border-left-color: var(--color-attention-emphasis);
  }
  
  .markdown-body .markdown-alert.markdown-alert-warning .markdown-alert-title {
	color: var(--color-attention-fg);
  }
  
  .markdown-body .markdown-alert.markdown-alert-tip {
	border-left-color: var(--color-success-emphasis);
  }
  
  .markdown-body .markdown-alert.markdown-alert-tip .markdown-alert-title {
	color: var(--color-success-fg);
  }
  
  .markdown-body .markdown-alert.markdown-alert-caution {
	border-left-color: var(--color-danger-emphasis);
  }
  
  .markdown-body .markdown-alert.markdown-alert-caution .markdown-alert-title {
	color: var(--color-danger-fg);
  }

  /*dark*/

.markdown-body {
  color-scheme: dark;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  margin: 0;
  color: #e6edf3;
  background-color: #0d1117;
  font-family: -apple-system,BlinkMacSystemFont,"Segoe UI","Noto Sans",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji";
  font-size: 16px;
  line-height: 1.5;
  word-wrap: break-word;
}

.markdown-body .octicon {
  display: inline-block;
  fill: currentColor;
  vertical-align: text-bottom;
}

.markdown-body h1:hover .anchor .octicon-link:before,
.markdown-body h2:hover .anchor .octicon-link:before,
.markdown-body h3:hover .anchor .octicon-link:before,
.markdown-body h4:hover .anchor .octicon-link:before,
.markdown-body h5:hover .anchor .octicon-link:before,
.markdown-body h6:hover .anchor .octicon-link:before {
  width: 16px;
  height: 16px;
  content: ' ';
  display: inline-block;
  background-color: currentColor;
  -webkit-mask-image: url("data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' version='1.1' aria-hidden='true'><path fill-rule='evenodd' d='M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z'></path></svg>");
  mask-image: url("data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' version='1.1' aria-hidden='true'><path fill-rule='evenodd' d='M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z'></path></svg>");
}

.markdown-body details,
.markdown-body figcaption,
.markdown-body figure {
  display: block;
}

.markdown-body summary {
  display: list-item;
}

.markdown-body [hidden] {
  display: none !important;
}

.markdown-body a {
  background-color: transparent;
  color: #2f81f7;
  text-decoration: none;
}

.markdown-body abbr[title] {
  border-bottom: none;
  -webkit-text-decoration: underline dotted;
  text-decoration: underline dotted;
}

.markdown-body b,
.markdown-body strong {
  font-weight: 600;
}

.markdown-body dfn {
  font-style: italic;
}

.markdown-body h1 {
  margin: .67em 0;
  font-weight: 600;
  padding-bottom: .3em;
  font-size: 2em;
  border-bottom: 1px solid #21262d;
}

.markdown-body mark {
  background-color: rgba(187,128,9,0.15);
  color: #e6edf3;
}

.markdown-body small {
  font-size: 90%;
}

.markdown-body sub,
.markdown-body sup {
  font-size: 75%;
  line-height: 0;
  position: relative;
  vertical-align: baseline;
}

.markdown-body sub {
  bottom: -0.25em;
}

.markdown-body sup {
  top: -0.5em;
}

.markdown-body img {
  border-style: none;
  max-width: 100%;
  box-sizing: content-box;
  background-color: #0d1117;
}

.markdown-body code,
.markdown-body kbd,
.markdown-body pre,
.markdown-body samp {
  font-family: monospace;
  font-size: 1em;
}

.markdown-body figure {
  margin: 1em 40px;
}

.markdown-body hr {
  box-sizing: content-box;
  overflow: hidden;
  background: transparent;
  border-bottom: 1px solid #21262d;
  height: .25em;
  padding: 0;
  margin: 24px 0;
  background-color: #30363d;
  border: 0;
}

.markdown-body input {
  font: inherit;
  margin: 0;
  overflow: visible;
  font-family: inherit;
  font-size: inherit;
  line-height: inherit;
}

.markdown-body [type=button],
.markdown-body [type=reset],
.markdown-body [type=submit] {
  -webkit-appearance: button;
  appearance: button;
}

.markdown-body [type=checkbox],
.markdown-body [type=radio] {
  box-sizing: border-box;
  padding: 0;
}

.markdown-body [type=number]::-webkit-inner-spin-button,
.markdown-body [type=number]::-webkit-outer-spin-button {
  height: auto;
}

.markdown-body [type=search]::-webkit-search-cancel-button,
.markdown-body [type=search]::-webkit-search-decoration {
  -webkit-appearance: none;
  appearance: none;
}

.markdown-body ::-webkit-input-placeholder {
  color: inherit;
  opacity: .54;
}

.markdown-body ::-webkit-file-upload-button {
  -webkit-appearance: button;
  appearance: button;
  font: inherit;
}

.markdown-body a:hover {
  text-decoration: underline;
}

.markdown-body ::placeholder {
  color: #6e7681;
  opacity: 1;
}

.markdown-body hr::before {
  display: table;
  content: "";
}

.markdown-body hr::after {
  display: table;
  clear: both;
  content: "";
}

.markdown-body table {
  border-spacing: 0;
  border-collapse: collapse;
  display: block;
  width: max-content;
  max-width: 100%;
  overflow: auto;
}

.markdown-body td,
.markdown-body th {
  padding: 0;
}

.markdown-body details summary {
  cursor: pointer;
}

.markdown-body details:not([open])>*:not(summary) {
  display: none !important;
}

.markdown-body a:focus,
.markdown-body [role=button]:focus,
.markdown-body input[type=radio]:focus,
.markdown-body input[type=checkbox]:focus {
  outline: 2px solid #2f81f7;
  outline-offset: -2px;
  box-shadow: none;
}

.markdown-body a:focus:not(:focus-visible),
.markdown-body [role=button]:focus:not(:focus-visible),
.markdown-body input[type=radio]:focus:not(:focus-visible),
.markdown-body input[type=checkbox]:focus:not(:focus-visible) {
  outline: solid 1px transparent;
}

.markdown-body a:focus-visible,
.markdown-body [role=button]:focus-visible,
.markdown-body input[type=radio]:focus-visible,
.markdown-body input[type=checkbox]:focus-visible {
  outline: 2px solid #2f81f7;
  outline-offset: -2px;
  box-shadow: none;
}

.markdown-body a:not([class]):focus,
.markdown-body a:not([class]):focus-visible,
.markdown-body input[type=radio]:focus,
.markdown-body input[type=radio]:focus-visible,
.markdown-body input[type=checkbox]:focus,
.markdown-body input[type=checkbox]:focus-visible {
  outline-offset: 0;
}

.markdown-body kbd {
  display: inline-block;
  padding: 3px 5px;
  font: 11px ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
  line-height: 10px;
  color: #e6edf3;
  vertical-align: middle;
  background-color: #161b22;
  border: solid 1px rgba(110,118,129,0.4);
  border-bottom-color: rgba(110,118,129,0.4);
  border-radius: 6px;
  box-shadow: inset 0 -1px 0 rgba(110,118,129,0.4);
}

.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
  margin-top: 24px;
  margin-bottom: 16px;
  font-weight: 600;
  line-height: 1.25;
}

.markdown-body h2 {
  font-weight: 600;
  padding-bottom: .3em;
  font-size: 1.5em;
  border-bottom: 1px solid #21262d;
}

.markdown-body h3 {
  font-weight: 600;
  font-size: 1.25em;
}

.markdown-body h4 {
  font-weight: 600;
  font-size: 1em;
}

.markdown-body h5 {
  font-weight: 600;
  font-size: .875em;
}

.markdown-body h6 {
  font-weight: 600;
  font-size: .85em;
  color: #848d97;
}

.markdown-body p {
  margin-top: 0;
  margin-bottom: 10px;
}

.markdown-body blockquote {
  margin: 0;
  padding: 0 1em;
  color: #848d97;
  border-left: .25em solid #30363d;
}

.markdown-body ul,
.markdown-body ol {
  margin-top: 0;
  margin-bottom: 0;
  padding-left: 2em;
}

.markdown-body ol ol,
.markdown-body ul ol {
  list-style-type: lower-roman;
}

.markdown-body ul ul ol,
.markdown-body ul ol ol,
.markdown-body ol ul ol,
.markdown-body ol ol ol {
  list-style-type: lower-alpha;
}

.markdown-body dd {
  margin-left: 0;
}

.markdown-body tt,
.markdown-body code,
.markdown-body samp {
  font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
  font-size: 12px;
}

.markdown-body pre {
  margin-top: 0;
  margin-bottom: 0;
  font-family: ui-monospace,SFMono-Regular,SF Mono,Menlo,Consolas,Liberation Mono,monospace;
  font-size: 12px;
  word-wrap: normal;
}

.markdown-body .octicon {
  display: inline-block;
  overflow: visible !important;
  vertical-align: text-bottom;
  fill: currentColor;
}

.markdown-body input::-webkit-outer-spin-button,
.markdown-body input::-webkit-inner-spin-button {
  margin: 0;
  -webkit-appearance: none;
  appearance: none;
}

.markdown-body .mr-2 {
  margin-right: 8px !important;
}

.markdown-body::before {
  display: table;
  content: "";
}

.markdown-body::after {
  display: table;
  clear: both;
  content: "";
}

.markdown-body>*:first-child {
  margin-top: 0 !important;
}

.markdown-body>*:last-child {
  margin-bottom: 0 !important;
}

.markdown-body a:not([href]) {
  color: inherit;
  text-decoration: none;
}

.markdown-body .absent {
  color: #f85149;
}

.markdown-body .anchor {
  float: left;
  padding-right: 4px;
  margin-left: -20px;
  line-height: 1;
}

.markdown-body .anchor:focus {
  outline: none;
}

.markdown-body p,
.markdown-body blockquote,
.markdown-body ul,
.markdown-body ol,
.markdown-body dl,
.markdown-body table,
.markdown-body pre,
.markdown-body details {
  margin-top: 0;
  margin-bottom: 16px;
}

.markdown-body blockquote>:first-child {
  margin-top: 0;
}

.markdown-body blockquote>:last-child {
  margin-bottom: 0;
}

.markdown-body h1 .octicon-link,
.markdown-body h2 .octicon-link,
.markdown-body h3 .octicon-link,
.markdown-body h4 .octicon-link,
.markdown-body h5 .octicon-link,
.markdown-body h6 .octicon-link {
  color: #e6edf3;
  vertical-align: middle;
  visibility: hidden;
}

.markdown-body h1:hover .anchor,
.markdown-body h2:hover .anchor,
.markdown-body h3:hover .anchor,
.markdown-body h4:hover .anchor,
.markdown-body h5:hover .anchor,
.markdown-body h6:hover .anchor {
  text-decoration: none;
}

.markdown-body h1:hover .anchor .octicon-link,
.markdown-body h2:hover .anchor .octicon-link,
.markdown-body h3:hover .anchor .octicon-link,
.markdown-body h4:hover .anchor .octicon-link,
.markdown-body h5:hover .anchor .octicon-link,
.markdown-body h6:hover .anchor .octicon-link {
  visibility: visible;
}

.markdown-body h1 tt,
.markdown-body h1 code,
.markdown-body h2 tt,
.markdown-body h2 code,
.markdown-body h3 tt,
.markdown-body h3 code,
.markdown-body h4 tt,
.markdown-body h4 code,
.markdown-body h5 tt,
.markdown-body h5 code,
.markdown-body h6 tt,
.markdown-body h6 code {
  padding: 0 .2em;
  font-size: inherit;
}

.markdown-body summary h1,
.markdown-body summary h2,
.markdown-body summary h3,
.markdown-body summary h4,
.markdown-body summary h5,
.markdown-body summary h6 {
  display: inline-block;
}

.markdown-body summary h1 .anchor,
.markdown-body summary h2 .anchor,
.markdown-body summary h3 .anchor,
.markdown-body summary h4 .anchor,
.markdown-body summary h5 .anchor,
.markdown-body summary h6 .anchor {
  margin-left: -40px;
}

.markdown-body summary h1,
.markdown-body summary h2 {
  padding-bottom: 0;
  border-bottom: 0;
}

.markdown-body ul.no-list,
.markdown-body ol.no-list {
  padding: 0;
  list-style-type: none;
}

.markdown-body ol[type="a s"] {
  list-style-type: lower-alpha;
}

.markdown-body ol[type="A s"] {
  list-style-type: upper-alpha;
}

.markdown-body ol[type="i s"] {
  list-style-type: lower-roman;
}

.markdown-body ol[type="I s"] {
  list-style-type: upper-roman;
}

.markdown-body ol[type="1"] {
  list-style-type: decimal;
}

.markdown-body div>ol:not([type]) {
  list-style-type: decimal;
}

.markdown-body ul ul,
.markdown-body ul ol,
.markdown-body ol ol,
.markdown-body ol ul {
  margin-top: 0;
  margin-bottom: 0;
}

.markdown-body li>p {
  margin-top: 16px;
}

.markdown-body li+li {
  margin-top: .25em;
}

.markdown-body dl {
  padding: 0;
}

.markdown-body dl dt {
  padding: 0;
  margin-top: 16px;
  font-size: 1em;
  font-style: italic;
  font-weight: 600;
}

.markdown-body dl dd {
  padding: 0 16px;
  margin-bottom: 16px;
}

.markdown-body table th {
  font-weight: 600;
}

.markdown-body table th,
.markdown-body table td {
  padding: 6px 13px;
  border: 1px solid #30363d;
}

.markdown-body table td>:last-child {
  margin-bottom: 0;
}

.markdown-body table tr {
  background-color: #0d1117;
  border-top: 1px solid #21262d;
}

.markdown-body table tr:nth-child(2n) {
  background-color: #161b22;
}

.markdown-body table img {
  background-color: transparent;
}

.markdown-body img[align=right] {
  padding-left: 20px;
}

.markdown-body img[align=left] {
  padding-right: 20px;
}

.markdown-body .emoji {
  max-width: none;
  vertical-align: text-top;
  background-color: transparent;
}

.markdown-body span.frame {
  display: block;
  overflow: hidden;
}

.markdown-body span.frame>span {
  display: block;
  float: left;
  width: auto;
  padding: 7px;
  margin: 13px 0 0;
  overflow: hidden;
  border: 1px solid #30363d;
}

.markdown-body span.frame span img {
  display: block;
  float: left;
}

.markdown-body span.frame span span {
  display: block;
  padding: 5px 0 0;
  clear: both;
  color: #e6edf3;
}

.markdown-body span.align-center {
  display: block;
  overflow: hidden;
  clear: both;
}

.markdown-body span.align-center>span {
  display: block;
  margin: 13px auto 0;
  overflow: hidden;
  text-align: center;
}

.markdown-body span.align-center span img {
  margin: 0 auto;
  text-align: center;
}

.markdown-body span.align-right {
  display: block;
  overflow: hidden;
  clear: both;
}

.markdown-body span.align-right>span {
  display: block;
  margin: 13px 0 0;
  overflow: hidden;
  text-align: right;
}

.markdown-body span.align-right span img {
  margin: 0;
  text-align: right;
}

.markdown-body span.float-left {
  display: block;
  float: left;
  margin-right: 13px;
  overflow: hidden;
}

.markdown-body span.float-left span {
  margin: 13px 0 0;
}

.markdown-body span.float-right {
  display: block;
  float: right;
  margin-left: 13px;
  overflow: hidden;
}

.markdown-body span.float-right>span {
  display: block;
  margin: 13px auto 0;
  overflow: hidden;
  text-align: right;
}

.markdown-body code,
.markdown-body tt {
  padding: .2em .4em;
  margin: 0;
  font-size: 85%;
  white-space: break-spaces;
  background-color: rgba(110,118,129,0.4);
  border-radius: 6px;
}

.markdown-body code br,
.markdown-body tt br {
  display: none;
}

.markdown-body del code {
  text-decoration: inherit;
}

.markdown-body samp {
  font-size: 85%;
}

.markdown-body pre code {
  font-size: 100%;
}

.markdown-body pre>code {
  padding: 0;
  margin: 0;
  word-break: normal;
  white-space: pre;
  background: transparent;
  border: 0;
}

.markdown-body .highlight {
  margin-bottom: 16px;
}

.markdown-body .highlight pre {
  margin-bottom: 0;
  word-break: normal;
}

.markdown-body .highlight pre,
.markdown-body pre {
  padding: 16px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  color: #e6edf3;
  background-color: #161b22;
  border-radius: 6px;
}

.markdown-body pre code,
.markdown-body pre tt {
  display: inline;
  max-width: auto;
  padding: 0;
  margin: 0;
  overflow: visible;
  line-height: inherit;
  word-wrap: normal;
  background-color: transparent;
  border: 0;
}

.markdown-body .csv-data td,
.markdown-body .csv-data th {
  padding: 5px;
  overflow: hidden;
  font-size: 12px;
  line-height: 1;
  text-align: left;
  white-space: nowrap;
}

.markdown-body .csv-data .blob-num {
  padding: 10px 8px 9px;
  text-align: right;
  background: #0d1117;
  border: 0;
}

.markdown-body .csv-data tr {
  border-top: 0;
}

.markdown-body .csv-data th {
  font-weight: 600;
  background: #161b22;
  border-top: 0;
}

.markdown-body [data-footnote-ref]::before {
  content: "[";
}

.markdown-body [data-footnote-ref]::after {
  content: "]";
}

.markdown-body .footnotes {
  font-size: 12px;
  color: #848d97;
  border-top: 1px solid #30363d;
}

.markdown-body .footnotes ol {
  padding-left: 16px;
}

.markdown-body .footnotes ol ul {
  display: inline-block;
  padding-left: 16px;
  margin-top: 16px;
}

.markdown-body .footnotes li {
  position: relative;
}

.markdown-body .footnotes li:target::before {
  position: absolute;
  top: -8px;
  right: -8px;
  bottom: -8px;
  left: -24px;
  pointer-events: none;
  content: "";
  border: 2px solid #1f6feb;
  border-radius: 6px;
}

.markdown-body .footnotes li:target {
  color: #e6edf3;
}

.markdown-body .footnotes .data-footnote-backref g-emoji {
  font-family: monospace;
}

.markdown-body .pl-c {
  color: #8b949e;
}

.markdown-body .pl-c1,
.markdown-body .pl-s .pl-v {
  color: #79c0ff;
}

.markdown-body .pl-e,
.markdown-body .pl-en {
  color: #d2a8ff;
}

.markdown-body .pl-smi,
.markdown-body .pl-s .pl-s1 {
  color: #c9d1d9;
}

.markdown-body .pl-ent {
  color: #7ee787;
}

.markdown-body .pl-k {
  color: #ff7b72;
}

.markdown-body .pl-s,
.markdown-body .pl-pds,
.markdown-body .pl-s .pl-pse .pl-s1,
.markdown-body .pl-sr,
.markdown-body .pl-sr .pl-cce,
.markdown-body .pl-sr .pl-sre,
.markdown-body .pl-sr .pl-sra {
  color: #a5d6ff;
}

.markdown-body .pl-v,
.markdown-body .pl-smw {
  color: #ffa657;
}

.markdown-body .pl-bu {
  color: #f85149;
}

.markdown-body .pl-ii {
  color: #f0f6fc;
  background-color: #8e1519;
}

.markdown-body .pl-c2 {
  color: #f0f6fc;
  background-color: #b62324;
}

.markdown-body .pl-sr .pl-cce {
  font-weight: bold;
  color: #7ee787;
}

.markdown-body .pl-ml {
  color: #f2cc60;
}

.markdown-body .pl-mh,
.markdown-body .pl-mh .pl-en,
.markdown-body .pl-ms {
  font-weight: bold;
  color: #1f6feb;
}

.markdown-body .pl-mi {
  font-style: italic;
  color: #c9d1d9;
}

.markdown-body .pl-mb {
  font-weight: bold;
  color: #c9d1d9;
}

.markdown-body .pl-md {
  color: #ffdcd7;
  background-color: #67060c;
}

.markdown-body .pl-mi1 {
  color: #aff5b4;
  background-color: #033a16;
}

.markdown-body .pl-mc {
  color: #ffdfb6;
  background-color: #5a1e02;
}

.markdown-body .pl-mi2 {
  color: #c9d1d9;
  background-color: #1158c7;
}

.markdown-body .pl-mdr {
  font-weight: bold;
  color: #d2a8ff;
}

.markdown-body .pl-ba {
  color: #8b949e;
}

.markdown-body .pl-sg {
  color: #484f58;
}

.markdown-body .pl-corl {
  text-decoration: underline;
  color: #a5d6ff;
}

.markdown-body g-emoji {
  display: inline-block;
  min-width: 1ch;
  font-family: "Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";
  font-size: 1em;
  font-style: normal !important;
  font-weight: 400;
  line-height: 1;
  vertical-align: -0.075em;
}

.markdown-body g-emoji img {
  width: 1em;
  height: 1em;
}

.markdown-body .task-list-item {
  list-style-type: none;
}

.markdown-body .task-list-item label {
  font-weight: 400;
}

.markdown-body .task-list-item.enabled label {
  cursor: pointer;
}

.markdown-body .task-list-item+.task-list-item {
  margin-top: 4px;
}

.markdown-body .task-list-item .handle {
  display: none;
}

.markdown-body .task-list-item-checkbox {
  margin: 0 .2em .25em -1.4em;
  vertical-align: middle;
}

.markdown-body .contains-task-list:dir(rtl) .task-list-item-checkbox {
  margin: 0 -1.6em .25em .2em;
}

.markdown-body .contains-task-list {
  position: relative;
}

.markdown-body .contains-task-list:hover .task-list-item-convert-container,
.markdown-body .contains-task-list:focus-within .task-list-item-convert-container {
  display: block;
  width: auto;
  height: 24px;
  overflow: visible;
  clip: auto;
}

.markdown-body ::-webkit-calendar-picker-indicator {
  filter: invert(50%);
}

.markdown-body .markdown-alert {
  padding: 8px 16px;
  margin-bottom: 16px;
  color: inherit;
  border-left: .25em solid #30363d;
}

.markdown-body .markdown-alert>:first-child {
  margin-top: 0;
}

.markdown-body .markdown-alert>:last-child {
  margin-bottom: 0;
}

.markdown-body .markdown-alert .markdown-alert-title {
  display: flex;
  font-weight: 500;
  align-items: center;
  line-height: 1;
}

.markdown-body .markdown-alert.markdown-alert-note {
  border-left-color: #1f6feb;
}

.markdown-body .markdown-alert.markdown-alert-note .markdown-alert-title {
  color: #2f81f7;
}

.markdown-body .markdown-alert.markdown-alert-important {
  border-left-color: #8957e5;
}

.markdown-body .markdown-alert.markdown-alert-important .markdown-alert-title {
  color: #a371f7;
}

.markdown-body .markdown-alert.markdown-alert-warning {
  border-left-color: #9e6a03;
}

.markdown-body .markdown-alert.markdown-alert-warning .markdown-alert-title {
  color: #d29922;
}

.markdown-body .markdown-alert.markdown-alert-tip {
  border-left-color: #238636;
}

.markdown-body .markdown-alert.markdown-alert-tip .markdown-alert-title {
  color: #3fb950;
}

.markdown-body .markdown-alert.markdown-alert-caution {
  border-left-color: #da3633;
}

.markdown-body .markdown-alert.markdown-alert-caution .markdown-alert-title {
  color: #f85149;
}
`
