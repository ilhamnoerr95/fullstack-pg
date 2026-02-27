<script setup lang="ts">
import { ref, onMounted } from "vue";
import { message } from "ant-design-vue";
import type { TableColumnsType } from "ant-design-vue";
import dayjs from "dayjs";

import PaymentSummary from "../../components/organisms/PaymentSummary.vue";
import AppHeader from "../../components/organisms/AppHeader.vue";

// type
import { IPayment, IResponsePayments, ISummary } from "../../types/payment";

// helpers
import { api } from "../../helpers/api";

const payments = ref<IPayment[]>([]);
const summary = ref<ISummary>({
	failed: 0,
	success: 0,
	total: 0,
});

const loading = ref(false);

const columns: TableColumnsType<IPayment> = [
	{
		title: "Payment ID",
		dataIndex: "id",
		key: "id",
	},
	{
		title: "Merchant Name",
		dataIndex: "merchant_name",
		key: "merchant_name",
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

async function fetchPayments() {
	try {
		loading.value = true;
		const token = (localStorage.getItem("token") as string) || "";

		const response = await api.get<IResponsePayments>("/payments", "v1", token);

		payments.value = response.payments;
		summary.value = response.summary;
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
			<PaymentSummary :summary="summary" />

			<!-- Table -->
			<a-table
				:columns="columns"
				:data-source="payments"
				:loading="loading"
				rowKey="id"
				bordered
			>
				<template #bodyCell="{ column, text }">
					<span v-if="column.dataIndex === 'date'">
						{{ dayjs(text).format("DD MMM YYYY, HH:mm") }}
					</span>
					<span v-if="column.dataIndex === 'status'">
						<a-tag :color="text === 'success' ? 'green' : 'red'">{{
							text
						}}</a-tag>
					</span>
				</template>
			</a-table>
		</a-layout-content>
	</a-layout>
</template>
