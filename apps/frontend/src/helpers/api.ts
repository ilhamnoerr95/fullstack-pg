import { fetcher } from "../lib/fetcher";

export const api = {
	get: <T>(url: string, version: string, token?: string) =>
		fetcher<T>(url, version, { method: "GET", token }),

	post: <T>(url: string, version: string, body: any, token?: string) =>
		fetcher<T>(url, version, { method: "POST", body, token }),

	put: <T>(url: string, version: string, body: any, token?: string) =>
		fetcher<T>(url, version, { method: "PUT", body, token }),

	delete: <T>(url: string, version: string, token?: string) =>
		fetcher<T>(url, version, { method: "DELETE", token }),
};
