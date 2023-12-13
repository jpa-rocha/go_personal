let activeIndex = 0;

const slides = document.getElementsByTagName("article");

const handleRightClick = () => {
  const nextIndex = activeIndex - 1 >= 0 ? activeIndex - 1 : slides.length - 1;
  
  const currentSlide = document.querySelector(`[data-index="${activeIndex}"]`),
        nextSlide = document.querySelector(`[data-index="${nextIndex}"]`);
        
  currentSlide.dataset.status = "after";
  
  nextSlide.dataset.status = "becoming-active-from-before";
  
  setTimeout(() => {
    nextSlide.dataset.status = "active";
    activeIndex = nextIndex;
  });
}

const handleLeftClick = () => {
  const nextIndex = activeIndex + 1 <= slides.length - 1 ? activeIndex + 1 : 0;
  
  const currentSlide = document.querySelector(`[data-index="${activeIndex}"]`),
        nextSlide = document.querySelector(`[data-index="${nextIndex}"]`);
  
  currentSlide.dataset.status = "before";
  
  nextSlide.dataset.status = "becoming-active-from-after";
  
  setTimeout(() => {
    nextSlide.dataset.status = "active";
    activeIndex = nextIndex;
  });
}

function hackText(element) {
    const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    document.getElementById(element.id).onmouseover = event => {
        console.log("HI!")
        let iterations = 0;
        const interval = setInterval(() => {
            event.target.innerText = event.target.innerText.split("").map((letter, index) => {
                if (index < iterations) {
                    return event.target.dataset.value[index];
                }
                return letters[Math.floor(Math.random() * 26)];
            }).join("");
            if(iterations >= event.target.dataset.value.length) clearInterval(interval)
            iterations += 1 / 3;
        }, 30);
    }
}

// const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// document.getElementsByClassName("article-title-section").onmouseover = event => {
//     console.log("HI!")
//     let iterations = 0;
//     const interval = setInterval(() => {
//         event.target.innerText = event.target.innerText.split("").map((letter, index) => {
//             if (index < iterations) {
//                 return event.target.dataset.value[index];
//             }
//             return letters[Math.floor(Math.random() * 26)];
//         }).join("");
//         if(iterations >= event.target.dataset.value.length) clearInterval(interval)
//         iterations += 1 / 3;
//     }, 30);
// }
