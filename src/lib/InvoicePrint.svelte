<script lang="ts">
    import { onMount } from "svelte";
    import { invoice } from "../store";
    import { DecimalFixed } from "./decimal";

    let from: Array<string> = [];
    let to: Array<string> = [];
    let notes: Array<string> = [];

    onMount(() => {
        from = $invoice.from.split("\n");
        to = $invoice.to.split("\n");
        notes = $invoice.notes.split("\n");

        setTimeout(() => {
            window.print();
        }, 1000);
    });
</script>

<div id="invoice-container">
    <header class="clearfix">
        <div id="logo">
            <img src={$invoice.logoBase64Img} alt={$invoice.logoFilename} />
        </div>
        <h1>{$invoice.invoiceNumber}</h1>
        <div id="company" class="clearfix">
            {#each from as f}
                <div>{f}</div>
            {/each}
        </div>
        <div id="project">
            <div><span>ISSUE DATE</span> {$invoice.issueDate}</div>
            <div><span>DUE DATE</span> {$invoice.dueDate}</div>
            <div>
                {#each to as t, i}
                    {#if i == 0}
                        <span>BILL TO</span> {t} <br />
                    {:else}
                        <span></span> {t} <br />
                    {/if}
                {/each}
            </div>
        </div>
    </header>

    <main>
        <table>
            <thead>
                <tr>
                    <!--<th class="service">SERVICE</th>-->
                    <th class="desc">DESCRIPTION</th>
                    <th>PRICE</th>
                    <th>QTY</th>
                    <th>TOTAL</th>
                </tr>
            </thead>
            <tbody>
                {#each $invoice.rows as row}
                    <tr>
                        <!--<td class="service">Development</td>-->
                        <td class="desc">{row.description}</td>
                        <td class="rate">{$invoice.currencySymbol}{row.rate}</td>
                        <td class="qty">{row.quantity}</td>
                        <td class="total">{$invoice.currencySymbol}{row.amount}</td>
                    </tr>
                {/each}
                <!--<tr>
          <td colspan="4">SUBTOTAL</td>
          <td class="total">$5,200.00</td>
        </tr>-->
                <!--<tr>
          <td colspan="4">TAX 25%</td>
          <td class="total">$1,300.00</td>
        </tr>-->
                <tr>
                    <td colspan="3" class="grand total">TOTAL</td>
                    <td class="grand total"
                        >{$invoice.currencySymbol}{DecimalFixed($invoice.totalAmount)}</td
                    >
                </tr>
            </tbody>
        </table>

        {#if notes.length > 0}
            <div id="notices">
                <div><b>NOTES</b></div> <br />
                <div class="notice">
                    {#each notes as note}
                        {note} <br />
                    {/each}
                </div>
            </div>
        {/if}
    </main>

    <footer>
        Invoice was created by <a href="https://goldenhandsoftware.co.uk"
            >Goldenhand Software Limited</a
        > and is valid without the signature.
    </footer>
</div>

<style>
    .clearfix:after {
        content: "";
        display: table;
        clear: both;
    }

    a {
        color: #5d6975;
        text-decoration: underline;
    }

    #invoice-container {
        position: relative;
        width: 21cm;
        height: 29.7cm;
        margin: 0 auto;
        color: #001028;
        background: #ffffff;
        font-family: Arial, sans-serif;
        font-size: 12px;
        font-family: Arial;
    }

    header {
        padding: 10px 0;
        margin-bottom: 30px;
    }

    #logo {
        text-align: center;
        margin-bottom: 10px;
    }

    #logo img {
        width: 90px;
    }

    h1 {
        border-top: 1px solid #5d6975;
        border-bottom: 1px solid #5d6975;
        color: #5d6975;
        font-size: 2.4em;
        line-height: 1.4em;
        font-weight: normal;
        text-align: center;
        margin: 0 0 20px 0;
    }

    #project {
        float: left;
    }

    #project span {
        color: #5d6975;
        text-align: right;
        width: 52px;
        margin-right: 10px;
        display: inline-block;
        font-size: 0.8em;
    }

    #company {
        float: right;
        text-align: right;
    }

    #project div,
    #company div {
        white-space: nowrap;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        border-spacing: 0;
        margin-bottom: 20px;
    }

    table tr:nth-child(2n-1) td {
        background: #f5f5f5;
    }

    table th,
    table td {
        text-align: center;
    }

    table th {
        padding: 5px 20px;
        color: #5d6975;
        border-bottom: 1px solid #c1ced9;
        white-space: nowrap;
        font-weight: normal;
        text-align: right;
    }

    /*table .service,*/
    table .desc {
        text-align: left;
    }

    table td {
        padding: 20px;
        text-align: right;
    }

    /*table td.service,*/
    table td.desc {
        vertical-align: top;
    }

    table td.rate,
    table td.qty,
    table td.total {
        font-size: 1.2em;
    }

    table td.grand {
        border-top: 1px solid #5d6975;
    }

    #notices .notice {
        color: #5d6975;
        font-size: 1.2em;
    }

    footer {
        color: #5d6975;
        width: 100%;
        height: 30px;
        position: absolute;
        bottom: 0;
        border-top: 1px solid #c1ced9;
        padding: 8px 0;
        text-align: center;
    }
</style>
