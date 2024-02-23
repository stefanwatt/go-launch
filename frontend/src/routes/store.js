import { writable, derived, get } from 'svelte/store';
import { mapSearchResults } from './desktop-entry';

//** @type {Writeable<App.DesktopEntry[]>}*/
export const desktopEntries = writable([]);

export const selectionPosition = writable({
	row: 0,
	col: 0
});

//** @type {Writeable<HTMLInputElement>}*/
export const keyboardNavigationInput = writable();

//** @type {Writeable<HTMLInputElement>}*/
export const promptInput = writable();

//** @type {Writeable<string>}*/
export const searchTerm = writable('');

//** @type {Derived<App.DesktopEntry[]>}*/
export const searchResults = derived(
	[searchTerm, desktopEntries],
	([$searchTerm, $desktopEntries]) => {
		return $searchTerm === '' ? [[], [], [], []] : mapSearchResults($desktopEntries, $searchTerm);
	}
);
export const selectedEntry = derived(selectionPosition, ($selectionPosition) => {
	const $searchResults = get(searchResults);
	if (!$selectionPosition || $searchResults.every((row) => row.length === 0)) return null;
	if ($searchResults[0].length === 1) return $searchResults[0][0];
	return $searchResults[$selectionPosition.row][$selectionPosition.col];
});
