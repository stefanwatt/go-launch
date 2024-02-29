// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
		interface DesktopEntry {
			Name: string;
			Exec: string;
			Icon: string;
			Comment: string;
		}
	}
}

export {};
