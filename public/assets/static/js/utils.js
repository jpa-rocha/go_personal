let activeIndex = 0;

const groups = document.getElementsByClassName("article-group")

const handleRightClick = () => {

    const nextIndex = activeIndex + 1 <= groups.length - 1 ? activeIndex + 1 : 0;

    const currentGroup = document.querySelector(`[data-index="${activeIndex}"]`),
        nextGroup = document.querySelector(`[data-index="${nextIndex}"]`);

    currentGroup.dataset.status = "inactive"
 
    nextGroup.dataset.status = "to-be-active-right"

    setTimeout(() => {

        nextGroup.dataset.status = "active"
        activeIndex = nextIndex;
    });

}

const handleLeftClick = () => {

    const nextIndex = activeIndex - 1 >= 0 ? activeIndex - 1 : groups.lenght - 1;

    const currentGroup = document.querySelector(`[data-index="${activeIndex}"]`),
        nextGroup = document.querySelector(`[data-index="${nextIndex}"]`);

    currentGroup.dataset.status = "inactive"
 
    nextGroup.dataset.status = "to-be-active-left"

    setTimeout(() => {

        nextGroup.dataset.status = "active"
        activeIndex = nextIndex;
    });

}
