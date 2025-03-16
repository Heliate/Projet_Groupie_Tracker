// Attendre que le DOM soit complètement chargé avant d'exécuter le code
document.addEventListener('DOMContentLoaded', function() {
    // Vérifier si le carrousel existe sur cette page avant de continuer
    const carousel = document.getElementById('boosterCarousel');
    if (!carousel) return;
    
    // Sélectionner les boutons de contrôle du carrousel
    const prevBtn = document.getElementById('prevBtn');
    const nextBtn = document.getElementById('nextBtn');
    const openBtn = document.getElementById('openBooster');
    
    // Configuration des boosters avec leurs informations et liens de redirection
    const boosterTypes = [
        { id: 1, name: "Sword & Shield", image: "/static/images/swsh1.png", link: "/cards?set=swsh1" },
        { id: 2, name: "Rebel Clash", image: "/static/images/swsh2.png", link: "/cards?set=swsh2" },
        { id: 3, name: "Darkness Ablaze", image: "/static/images/swsh3.png", link: "/cards?set=swsh3" },
        { id: 4, name: "Vivid Voltage", image: "/static/images/swsh4.png", link: "/cards?set=swsh4" },
        { id: 5, name: "Battle Styles", image: "/static/images/swsh5.png", link: "/cards?set=swsh5" },
        { id: 6, name: "Chilling Reign", image: "/static/images/swsh6.png", link: "/cards?set=swsh6" },
        { id: 7, name: "Evolving Skies", image: "/static/images/swsh7.png", link: "/cards?set=swsh7" },
        { id: 8, name: "Fusion Strike", image: "/static/images/swsh8.png", link: "/cards?set=swsh8" },
    ];
    
    // Variables de contrôle du carrousel
    let currentIndex = 0;         // Index du booster actuellement affiché
    let radius = 400;             // Rayon du cercle pour l'effet 3D
    let animating = false;        // Blocage pendant les animations
    
    // Fonction pour créer les éléments visuels du carrousel
    function createCarouselItems() {
        // Vider le carousel pour éviter les doublons
        carousel.innerHTML = '';
        
        // Créer un élément pour chaque booster
        boosterTypes.forEach((booster, index) => {
            const item = document.createElement('div');
            item.className = 'carousel-item';
            item.dataset.id = booster.id;
            
            // Ajouter l'image du booster
            const img = document.createElement('img');
            img.src = booster.image;
            img.alt = booster.name;
            
            item.appendChild(img);
            carousel.appendChild(item);
        });
        
        // Positionner tous les éléments en cercle
        updateCarousel();
        
        // Afficher le bouton d'ouverture une fois le carrousel chargé
        if (openBtn) openBtn.style.display = 'block';
    }
    
    // Fonction pour mettre à jour la position des éléments dans le carrousel
    function updateCarousel() {
        const items = document.querySelectorAll('.carousel-item');
        const angleStep = 360 / items.length;
        
        items.forEach((item, index) => {
            // Calculer la position angulaire de chaque élément
            const angle = angleStep * ((index - currentIndex + items.length) % items.length);
            const radians = (angle * Math.PI) / 180;
            
            // Calculer les coordonnées 3D sur le cercle
            const x = radius * Math.sin(radians);
            const z = radius * Math.cos(radians) - radius;
            
            // Appliquer la transformation 3D
            item.style.transform = `translate(calc(-50% + ${x}px), -50%) translateZ(${z}px)`;
            
            // Ajuster l'opacité selon la position (plus transparent à l'arrière)
            const opacity = Math.max(0.3, 1 - Math.abs(angle) / 180);
            item.style.opacity = opacity;
            
            // Ajuster le z-index pour que les éléments au premier plan soient au-dessus
            const zIndex = Math.round(100 - Math.abs(angle));
            item.style.zIndex = zIndex;
            
            // Marquer l'élément actuellement sélectionné (à l'avant)
            if (index === currentIndex) {
                item.classList.add('selected');
            } else {
                item.classList.remove('selected');
            }
        });
    }
    
    // Fonction pour faire tourner le carrousel vers un index spécifique
    function rotateToIndex(index) {
        if (animating || index === currentIndex) return;
        
        animating = true;
        currentIndex = index;
        updateCarousel();
        
        // Débloquer l'animation après la transition
        setTimeout(() => {
            animating = false;
        }, 500); // Durée correspondant à la transition CSS
    }
    
    // Fonction pour passer au booster suivant
    function rotateNext() {
        if (animating) return;
        const nextIndex = (currentIndex + 1) % boosterTypes.length;
        rotateToIndex(nextIndex);
    }
    
    // Fonction pour revenir au booster précédent
    function rotatePrev() {
        if (animating) return;
        const prevIndex = (currentIndex - 1 + boosterTypes.length) % boosterTypes.length;
        rotateToIndex(prevIndex);
    }
    
    // Fonction pour "ouvrir" le booster (avec animation et redirection)
    function openBooster() {
        if (animating) return;
        
        const selectedItem = document.querySelector('.carousel-item.selected');
        if (selectedItem) {
            animating = true;
            selectedItem.classList.add('opening-animation');
            
            setTimeout(() => {
                selectedItem.classList.remove('opening-animation');
                animating = false;
                
                // Récupérer les informations du booster sélectionné
                const boosterId = parseInt(selectedItem.dataset.id);
                const booster = boosterTypes.find(b => b.id === boosterId);
                
                // Rediriger vers la page des cartes du set correspondant
                if (booster && booster.link) {
                    window.location.href = booster.link;
                }
                
            }, 1000); // Durée de l'animation d'ouverture
        }
    }
    
    // Configuration des gestionnaires d'événements pour les boutons
    if (prevBtn) prevBtn.addEventListener('click', rotatePrev);
    if (nextBtn) nextBtn.addEventListener('click', rotateNext);
    if (openBtn) openBtn.addEventListener('click', openBooster);
    
    // Support des touches du clavier pour la navigation
    document.addEventListener('keydown', function(e) {
        if (e.key === 'ArrowLeft') {
            rotatePrev();
        } else if (e.key === 'ArrowRight') {
            rotateNext();
        } else if (e.key === 'Enter') {
            openBooster();
        }
    });
    
    // Support des gestes tactiles pour la navigation sur mobile
    let touchStartX = 0;
    let touchEndX = 0;
    
    // Détecter le début du toucher
    carousel.addEventListener('touchstart', function(e) {
        touchStartX = e.changedTouches[0].screenX;
    });
    
    // Détecter la fin du toucher et traiter le geste
    carousel.addEventListener('touchend', function(e) {
        touchEndX = e.changedTouches[0].screenX;
        handleSwipe();
    });
    
    // Interprétation du geste de balayage
    function handleSwipe() {
        const minSwipeDistance = 50; // Distance minimum pour considérer un balayage valide
        if (touchEndX < touchStartX - minSwipeDistance) {
            // Balayage vers la gauche -> booster suivant
            rotateNext();
        } else if (touchEndX > touchStartX + minSwipeDistance) {
            // Balayage vers la droite -> booster précédent
            rotatePrev();
        }
    }
    
    // Initialiser le carrousel au chargement
    createCarouselItems();
    
    // Fonction pour ajuster le rayon en fonction de la taille de l'écran
    function adjustRadius() {
        const containerWidth = document.querySelector('.carousel').clientWidth;
        radius = Math.min(400, containerWidth * 0.45); // Limiter le rayon sur les petits écrans
        updateCarousel();
    }
    
    // Ajuster le rayon lors du redimensionnement de la fenêtre
    window.addEventListener('resize', adjustRadius);
    
    // Ajuster le rayon une première fois au chargement
    adjustRadius();
});