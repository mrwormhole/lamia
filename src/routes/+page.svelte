<script lang="ts">
    import Header from "$lib/Header.svelte";
    import Footer from "$lib/Footer.svelte";
    import InvoiceForm from "$lib/InvoiceForm.svelte";
</script>

<div id="main">
    <Header appName="Lamia Invoice Generator" />
    <InvoiceForm />
    <Footer appName="Lamia Invoice Generator" />
</div>
