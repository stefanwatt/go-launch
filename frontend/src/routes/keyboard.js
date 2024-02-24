import { LaunchApp, HideLauncher } from '$lib/wailsjs/go/main/App';
import { get } from 'svelte/store';
import {
	selectionPosition,
	searchTerm as searchTerm$,
	searchResults as searchResults$,
	selectedEntry as selectedEntry$,
	promptInput
} from './store';

/** @param {HTMLInputElement} node */
export function onKeyPress(node) {
	let ctrlDown = false;
	window.addEventListener('keydown', (event) => {
		if (event.key === 'Control') return (ctrlDown = true);
		if (event.key === 'k' && ctrlDown) {
			event.preventDefault();
			selectionPosition.set(null);
			const input = get(promptInput);
			input.focus();
		}
	});
	window.addEventListener('keyup', (event) => {
		if (event.key === 'Control') ctrlDown = false;
	});
	node.addEventListener('keyup', (event) => {
		const searchResults = get(searchResults$);
		const searchTerm = get(searchTerm$);
		const selectedEntry = get(selectedEntry$);
		if (!selectedEntry) return;
		const row = searchResults.findIndex((row) => row.includes(selectedEntry));
		const col = searchResults[row].indexOf(selectedEntry);

		let nextRow, nextCol;
		const size = 4;
		switch (event.key) {
			case 'ArrowUp':
				if (row === 0) return selectionPosition.set(null);
				selectionPosition.set({ row: row - 1, col });
				break;
			case 'ArrowDown':
				nextRow = row + 1;
				if (!(searchResults[nextRow]?.length - 1 >= col))
					return selectionPosition.set({ row: 0, col });
				selectionPosition.set({ row: nextRow, col });
				break;
			case 'ArrowLeft':
				nextCol = col - 1;
				if (searchResults[row].length === 1 && row > 0)
					return selectionPosition.set({ row: row - 1, col: size - 1 });
				if (nextCol < 0) return selectionPosition.set({ row, col: size - 1 });
				selectionPosition.set({ row, col: nextCol });
				break;
			case 'ArrowRight':
				nextCol = col + 1;
				if (!searchResults[row + 1]?.length && !searchResults[row][nextCol])
					return selectionPosition.set({ row: 0, col: 0 });
				if (nextCol >= size) return selectionPosition.set({ row, col: 0 });
				selectionPosition.set({ row, col: nextCol });
				break;
			case 'Enter':
				LaunchApp(selectedEntry.Exec);
				searchTerm$.set('');
				break;
			case 'Escape':
				HideLauncher();
				searchTerm$.set('');
				break;
			case 'Backspace':
				event.preventDefault();
				selectionPosition.set(null);
				const input = get(promptInput);
				input.focus();
				searchTerm$.set(searchTerm.slice(0, -1));
				break;
			default:
				break;
		}
	});
}

/** @param {KeyboardEvent} e*/
export function onInputKeypress(e) {
	if (e.key === 'Escape') HideLauncher();
	if (get(selectedEntry$)) return;
	if (e.key === 'ArrowDown') selectionPosition.set({ row: 0, col: 0 });
}
