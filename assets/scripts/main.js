$(document).ready(function () {
    console.log("Script linked successfully.");
    console.log(mdc);

    const topAppBarElement = document.querySelector('.mdc-top-app-bar');
    const topAppBar = mdc.topAppBar.MDCTopAppBar.attachTo(topAppBarElement);
    const drawer = mdc.drawer.MDCDrawer.attachTo(document.querySelector('.mdc-drawer'));

    topAppBar.setScrollTarget(document.getElementById('scrollbar'));
    topAppBar.listen('MDCTopAppBar:nav', () => {
        drawer.open = !drawer.open;
    });

    // const listEl = document.querySelector('.mdc-drawer .mdc-list');
    // const mainContentEl = document.querySelector('.mdc-drawer-scrim');
    // listEl.addEventListener('click', (event) => {
    //     drawer.open = false;
    // });
    // document.body.addEventListener('MDCDrawer:closed', () => {
    //     mainContentEl.querySelector('input, button').focus();
    // });

    if ($('#sessionUser').text() != "") {
        const menu = mdc.menu.MDCMenu.attachTo(document.querySelector('.mdc-menu'));
        const dashButton = document.querySelector('#dashboard');

        dashButton.addEventListener('click', () => {
            menu.open = !menu.open;
        });
    }
});