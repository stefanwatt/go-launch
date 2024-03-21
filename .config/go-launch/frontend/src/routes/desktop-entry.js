/**
 * @param {DesktopEntry[]} desktopEntries
 * @param {string} searchTerm
 * @returns {DesktopEntry[][]}
 * */
export function mapSearchResults(desktopEntries, searchTerm) {
	const filteredEntries =
		searchTerm === ''
			? []
			: desktopEntries.filter((entry) =>
					entry.Name.toLowerCase().includes(searchTerm.toLowerCase())
				);
	const results = [[], [], [], []];
	const cols = 4;
	const rows = 4;
	for (let i = 0; i < rows; i++) {
		for (let j = 0; j < cols; j++) {
			const index = i * rows + j;
			if (filteredEntries[index]) {
				results[i][j] = filteredEntries[index];
			}
		}
	}
	return results;
}
