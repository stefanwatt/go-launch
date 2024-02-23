import { LaunchApp, HideLauncher } from '$lib/wailsjs/go/main/App';
import { get } from 'svelte/store';
import {
	selectionPosition,
	searchTerm as searchTerm$,
	searchResults as searchResults$,
	selectedEntry as selectedEntry$
} from './store';

const defaultPosition = { row: 0, col: 0 };

/**
 * @param {KeyboardEvent} event
 * */
export function onKeyPress(event) {
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
			console.log('ArrowUp');
			nextRow = row - 1;
			if (nextRow < 0 || nextRow >= size) return selectionPosition.set(defaultPosition);
			selectionPosition.set({ row: nextRow, col });
			break;
		case 'ArrowDown':
			console.log('ArrowDown');
			nextRow = row + 1;
			if (nextRow < 0 || nextRow >= size) return selectionPosition.set(defaultPosition);
			selectionPosition.set({ row: nextRow, col });
			break;
		case 'ArrowLeft':
			console.log('ArrowLeft');
			nextCol = col - 1;
			if (nextCol < 0 || nextCol >= size) return selectionPosition.set(defaultPosition);
			selectionPosition.set({ row, col: nextCol });
			break;
		case 'ArrowRight':
			console.log('ArrowRight');
			nextCol = col + 1;
			if (nextCol < 0 || nextCol >= size) return selectionPosition.set(defaultPosition);
			selectionPosition.set({ row, col: nextCol });
			break;
		case 'Enter':
			LaunchApp(selectedEntry.Exec);
			searchTerm = '';
			break;
		case 'Escape':
			searchTerm = '';
			HideLauncher();
			break;
		default:
			break;
	}
}

/** @param {KeyboardEvent} e*/
export function onInputKeypress(e) {
	if (e.key === 'Escape') HideLauncher();
	if (get(selectedEntry$)) return;
}
