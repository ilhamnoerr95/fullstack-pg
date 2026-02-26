import { useAuthStore } from "../store/auth.store";

export function authGuard() {
	const auth = useAuthStore();

	// if not authenticated, redirect to login page
	if (!auth.isAuthenticated) {
		console.log("Not authenticated, redirecting to login");
		return { name: "login" };
	}
	console.log("Authenticated, access granted");
}
