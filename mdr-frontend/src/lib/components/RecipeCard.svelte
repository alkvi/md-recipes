<script>
    export const content = "test";
</script>

<div class="main">
    <h1>
        <slot name="card_title"/>
    </h1>
    <slot/>
</div>

<style>

    .main {
        border: 2px solid #3e3939;
        border-radius: 4px;
        padding: .5rem;
    }

</style>
