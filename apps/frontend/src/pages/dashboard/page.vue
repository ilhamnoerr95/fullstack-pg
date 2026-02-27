<script setup lang="ts">
import { ref, onMounted } from "vue";
import { message } from "ant-design-vue";
import type { TableColumnsType } from "ant-design-vue";

import PaymentSummary from "../../components/organisms/PaymentSummary.vue";
import AppHeader from "../../components/organisms/AppHeader.vue";

// type
import { IPayment } from "../../types/payment";

const payments = ref<IPayment[]>([]);
const loading = ref(false);

const columns: TableColumnsType<IPayment> = [
	{
		title: "Payment ID",
		dataIndex: "id",
		key: "id",
	},
	{
		title: "Merchant Name",
		dataIndex: "merchantName",
		key: "merchantName",
	},
	{
		title: "Date",
		dataIndex: "date",
		key: "date",
	},
	{
		title: "Amount",
		dataIndex: "amount",
		key: "amount",
		customRender: ({ text }: any) => `Rp ${text.toLocaleString()}`,
	},
	{
		title: "Status",
		dataIndex: "status",
		key: "status",
	},
];

// 🔥 Simulasi API call
async function fetchPayments() {
	try {
		loading.value = true;

		// ganti dengan API asli nanti
		await new Promise((r) => setTimeout(r, 800));

		payments.value = [
			{
				id: 1,
				merchantName: "Tokopedia",
				date: "2026-02-25",
				amount: 150000,
				status: "SUCCESS",
			},
			{
				id: 2,
				merchantName: "Shopee",
				date: "2026-02-26",
				amount: 80000,
				status: "FAILED",
			},
			{
				id: 3,
				merchantName: "Bukalapak",
				date: "2026-02-26",
				amount: 250000,
				status: "SUCCESS",
			},
		];
	} catch (err) {
		message.error("Failed to fetch payments");
	} finally {
		loading.value = false;
	}
}

onMounted(fetchPayments);
</script>

<template>
	<a-layout style="background-color: #f0f2f5; min-height: 90vh">
		<AppHeader />
		<a-layout-content style="padding: 24px">
			<h1 style="margin-bottom: 20px; color: #00949e">Dashboard</h1>

			<!-- Summary -->
			<PaymentSummary :payments="payments" />

			<!-- Table -->
			<a-table
				:columns="columns"
				:data-source="payments"
				:loading="loading"
				rowKey="id"
				bordered
			>
				<template #bodyCell="{ column, text }">
					<span v-if="column.dataIndex === 'status'">
						<a-tag :color="text === 'SUCCESS' ? 'green' : 'red'">{{
							text
						}}</a-tag>
					</span>
				</template>
			</a-table>
		</a-layout-content>
	</a-layout>
</template>
