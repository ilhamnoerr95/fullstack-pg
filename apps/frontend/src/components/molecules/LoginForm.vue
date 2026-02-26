<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../../store/auth.store";

import BaseInput from "../atom/BaseInput.vue";
import GlassCard from "../atom/GlassCard.vue";

const auth = useAuthStore();
const router = useRouter();

const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const submit = async () => {
	error.value = "";
	loading.value = true;

	try {
		// fake API delay
		await new Promise((r) => setTimeout(r, 800));

		if (!email.value || !password.value) {
			throw new Error("Email dan password wajib diisi");
		}

		auth.login("fake-token", {
			id: 1,
			name: "Ilham",
			email: email.value,
		});

		router.push("/dashboard");
	} catch (e: any) {
		error.value = e.message;
	} finally {
		loading.value = false;
	}
};
</script>

<template>
	<form @submit.prevent="submit">
		<GlassCard>
			<h2>Login to dashboard</h2>

			<BaseInput
				v-model="email"
				label="Email"
				placeholder="you@email.com"
			/>

			<BaseInput
				v-model="password"
				label="Password"
				type="password"
				placeholder="••••••••"
			/>

			<p
				v-if="error"
				class="form-error"
			>
				{{ error }}
			</p>

			<button :disabled="loading">
				{{ loading ? "Logging in..." : "Login" }}
			</button>
		</GlassCard>
	</form>
</template>

<style scoped>
h2 {
	margin-bottom: 2rem;
	text-align: left;
	color: #00949e;
}

button {
	width: 100%;
	padding: 10px;
	border-radius: 8px;
	border: none;
	background: #fbc037;
	color: black;
	font-weight: 600;
	cursor: pointer;
}
button:hover:not(:disabled) {
	background: #f9a825;
}

button:disabled {
	opacity: 0.6;
	cursor: not-allowed;
}

.form-error {
	color: #e74c3c;
	font-size: 14px;
	text-align: center;
}
</style>
