var menuSource = document.getElementById('menu-template').innerHTML;
var menuTemplate = Handlebars.compile(menuSource);
var menuData = {
  project: [
        { name: "grocerygod", platform: "Android",
          desc: "No idea what to eat? No problem. grocerygod is a meal generating app to help people plan healthy meals for up to a week. Sound useful? Have an Android? Check it out:",
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
        { name: "World Wide Reed", platform: "Go / Handlebars",
          desc: "This is my playground for learning web stuff. Those poor, poor <div>'s. Right now it's hacked together with some Bootstrap 3 and Handlebars, running on a go server hosted by Heroku. There's your serving of buzzwords for the day.",
          links: [
            { url: "https://github.com/rdallman/worldwidereed",
              img: "assets/dev/source.png"
            }
          ],
          pictures: [
            'assets/dev/gopher.png',
            'assets/dev/bootstrap.png',
            'assets/dev/handlebars.jpg',
            'assets/dev/heroku.png',
          ]
        }
    ]
};
document.getElementById('menu-placeholder').innerHTML = menuTemplate(menuData);
