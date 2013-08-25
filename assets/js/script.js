var menuSource = document.getElementById('menu-template').innerHTML;
var menuTemplate = Handlebars.compile(menuSource);
var menuData = {
  project: [
        { name: "grocerygod", platform: "Android",
          desc: "No idea what to eat? No problem. grocerygod is a meal generating app to help people plan healthy meals for up to a week. This started as a school project and then turned business. Have an Android? Check it out:",
          links: [
            { url: "https://play.google.com/store/apps/details?id=com.bet.grocerygod",
              img: "https://developer.android.com/images/brand/en_generic_rgb_wo_45.png"
            }
          ],
          pictures: [
            'assets/grocerygod/icon.png',
            'assets/grocerygod/home.png',
            'assets/grocerygod/nutrition.png',
            'assets/grocerygod/meals.png',
            'assets/grocerygod/shopping.png',
            'assets/grocerygod/meal.png'
          ]
        },
        { name: "This site", platform: "Go / Handlebars",
          desc: "fill me in",
          pictures: []
        },
    ]
};
document.getElementById('menu-placeholder').innerHTML = menuTemplate(menuData);
