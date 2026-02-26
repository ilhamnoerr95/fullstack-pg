import { createRouter, createWebHistory } from "vue-router";
import { authGuard } from "./guard";

export const router = createRouter({
	history: createWebHistory(),
	routes: [
		// 🌍 PUBLIC
		{
			path: "/login",
			name: "login",
			component: () => import("../pages/auth/page.vue"),
		},
		{
			path: "/dashboard",
			name: "dashboard",
			beforeEnter: authGuard,
			component: () => import("../pages/dashboard/page.vue"),
		},
		// 🏠 DEFAULT
		{
			path: "/",
			redirect: "/dashboard",
		},
	],
});

export default router;
