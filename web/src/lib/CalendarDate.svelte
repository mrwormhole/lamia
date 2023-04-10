<script lang="ts">
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";

    type DatePickEvent = {
        "message": {
            date: string
        }
    }

    const dispatch = createEventDispatcher<DatePickEvent>();

    export let id: string = "calendar-date";
    const title: string = capitalize(id);

    function capitalize(id: string): string {
        const parts = id
            .toLowerCase()
            .split("-")
            .map((part) => {
                return part.charAt(0).toUpperCase() + part.slice(1);
            });
        return parts.join(" ");
    }

    async function createCalendar() {
        const bc = await import("bulma-calendar");

        bc.default.attach(`#${id}`, {
            dateFormat: "dd/MM/yyyy",
            color: "dark",
            weekStart: 1,
        });

        // To access to bulmaCalendar instance of an element
        let element: any = document.querySelector(`#${id}`);
        if (element != undefined && element.bulmaCalendar != undefined) {
            element.bulmaCalendar.on("select", function (datepicker: any) {
                if (datepicker != undefined && datepicker.data != undefined) {                    
                    dispatch("message", {
                        date: datepicker.data.value(),
                    });
                }
            });
        }
    }

    onMount(() => {
        createCalendar();
    });
</script>

<p class="date-title"><b>{title}</b></p>
<input type="date" id={id} />

<style>
    @import "../../node_modules/bulma-calendar/dist/css/bulma-calendar.min.css";

    .date-title {
        font-size: 1rem;
        padding: 1rem 0rem;
    }
</style>
