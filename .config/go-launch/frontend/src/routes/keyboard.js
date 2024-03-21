import { LaunchApp } from '$lib/go-launch/app.iframe';
import { get } from 'svelte/store';
import { findLastIndex } from './utils';
import {
	selectionPosition,
	searchTerm as searchTerm$,
	searchResults as searchResults$,
	selectedEntry as selectedEntry$,
	promptInput
} from './store';
/** @param {KeyboardEvent} event */
export function onKeyPress(event) {
	/**@type {app.DesktopEntry[][]} */
	const searchResults = get(searchResults$);
	/**@type {string} */
	const searchTerm = get(searchTerm$);
	/**@type {app.DesktopEntry} */
	let selectedEntry = get(selectedEntry$);
	if (!selectedEntry) selectionPosition.set({ row: 0, col: 0 });
	selectedEntry = get(selectedEntry$);
	const row = searchResults.findIndex((row) => row.includes(selectedEntry));
	if (row === -1) return;
	const col = searchResults[row].indexOf(selectedEntry);

	/**@type {number} */
	const cursorPos = get(promptInput).selectionStart;
	let nextRow, nextCol;
	const size = 4;
	switch (event.key) {
		case 'ArrowUp':
			if (row !== 0) return selectionPosition.set({ row: row - 1, col });
			nextRow = findLastIndex(searchResults, (row) => row[col]);
			if (nextRow === -1) return;
			selectionPosition.set({ row: nextRow, col });
			break;
		case 'ArrowDown':
			nextRow = row + 1;
			if (!searchResults[nextRow] || !searchResults[nextRow][col]) {
				return selectionPosition.set({ row: 0, col });
			}
			selectionPosition.set({ row: nextRow, col });
			break;
		case 'ArrowLeft':
			if (cursorPos !== 0) return;
			nextCol = col - 1;
			const entriesInCol = searchResults[row].filter(Boolean).length;
			if (nextCol < 0) return selectionPosition.set({ row, col: entriesInCol - 1 });
			selectionPosition.set({ row, col: nextCol });
			break;
		case 'ArrowRight':
			if (cursorPos !== searchTerm.length) return;
			nextCol = col + 1;
			if (!searchResults[row + 1]?.length && !searchResults[row][nextCol])
				return selectionPosition.set({ row: 0, col: 0 });
			if (nextCol >= size) return selectionPosition.set({ row, col: 0 });
			selectionPosition.set({ row, col: nextCol });
			break;
		case 'Enter':
			LaunchApp(selectedEntry.Id);
			searchTerm$.set('');
			break;
		case 'Escape':
			// HideLauncher();
			searchTerm$.set('');
			break;
		default:
			break;
	}
}

// 	const characters = [
// 		'a',
// 		'b',
// 		'c',
// 		'd',
// 		'e',
// 		'f',
// 		'g',
// 		'h',
// 		'i',
// 		'j',
// 		'k',
// 		'l',
// 		'm',
// 		'n',
// 		'o',
// 		'p',
// 		'q',
// 		'r',
// 		's',
// 		't',
// 		'u',
// 		'v',
// 		'w',
// 		'x',
// 		'y',
// 		'z',
// 		'0',
// 		'1',
// 		'2',
// 		'3',
// 		'4',
// 		'5',
// 		'6',
// 		'7',
// 		'8',
// 		'9',
// 		'!',
// 		'"',
// 		'#',
// 		'$',
// 		'%',
// 		'&',
// 		"'",
// 		'(',
// 		')',
// 		'*',
// 		'+',
// 		',',
// 		'-',
// 		'.',
// 		'/',
// 		':',
// 		';',
// 		'<',
// 		'=',
// 		'>',
// 		'?',
// 		'@',
// 		'[',
// 		'\\',
// 		']',
// 		'^',
// 		'_',
// 		'`',
// 		'{',
// 		'|',
// 		'}',
// 		'~'
// 	];
// 	// addEventListener(characters, node, 'keyup', () => {
// 	// 	selectionPosition.set(null);
// 	// 	const input = get(promptInput);
// 	// 	input.focus();
// 	// 	searchTerm$.update((old) => old + event.key);
// 	// });
// }
//
// // function that takes a list of keys, node and event and adds event listener
// /** @param {string[]} keys
//  * @param {HTMLElement} node
//  * @param {string} event
//  * @param{function} cb
//  * @param {any[]}args
//  */
// function addEventListener(keys, node, event, cb, ...args) {
// 	keys.forEach((key) => {
// 		node.addEventListener(event, (e) => {
// 			e.preventDefault();
// 			if (e.key === key) {
// 				cb(args);
// 			}
// 		});
// 	});
// }
