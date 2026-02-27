const BASE_URL = import.meta.env.VITE_API_URL;

const apiProxy: Record<string, string> = {
	v1: "/dashboard/v1",
};

type FetchOptions = {
	method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
	body?: any;
	headers?: Record<string, string>;
	token?: string;
};

export async function fetcher<T = any>(
	endpoint: string,
	version?: string,
	options: FetchOptions = {},
): Promise<T> {
	const { method = "GET", body, headers = {}, token } = options;

	const url = `${apiProxy[version || "v1"]}${endpoint}`;

	const res = await fetch(url, {
		method,
		headers: {
			"Content-Type": "application/json",
			...(token && { Authorization: `Bearer ${token}` }),
			...headers,
		},
		...(body && { body: JSON.stringify(body) }),
	});

	const contentType = res.headers.get("content-type");

	let data: any;
	if (contentType?.includes("application/json")) {
		data = await res.json();
	} else {
		data = await res.text();
	}

	if (!res.ok) {
		throw new Error(data?.error || "Something went wrong");
	}

	return data;
}
