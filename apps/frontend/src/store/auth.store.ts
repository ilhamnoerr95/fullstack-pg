import { defineStore } from "pinia";

interface User {
	id: number;
	email: string;
	password: string;
	role?: "cs" | "operation" | "admin";
}

export const useAuthStore = defineStore("auth", {
	state: () => ({
		token: (localStorage.getItem("token") as string) || null,
		user: null as User | null,
	}),
	getters: {
		isAuthenticated: (state) => !!state.token,
		role: (state) => state.user?.role || null,
	},
	actions: {
		login(token: string, user: User) {
			this.token = token;
			this.user = user;
			localStorage.setItem("token", token);
		},
		logout() {
			this.token = null;
			this.user = null;
			localStorage.removeItem("token");
		},
	},
});
