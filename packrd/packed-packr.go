// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "a84972d22884b939fe3f43f3215d970b"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"12a6b0a1bf0120c804d3cebd22ffe6e9": "1f8b08000000000000ffcc574d6fdb38133e8bbf622aa08154e8958bf7e85d2fb04d9ad6d94d52c4edf650145e851ac94c645225292781a1ffbe20a9cfc8299aee650304923833cf7c907c663c9bc1b1481172e428138d295c3f8016a2502032285855b0bf15c0c9255c5c7e84b727cb8f3129137a9be40874936842d8b6145243403c9f0aaef15efbc4f3b3ad7d70b48f8a332a52f409f16e14f839d39bea3aa6623b2b0ac5ee67efde5ccd3e6956a8d9d9ca275e2e4b0a7e2e445e609c8b22e1792c643e33eb3e0909d10f25c239ea8d48d5fb84a70582d2b2a21af6c43bad8ac2c9cc22e339f11887f68f718d324b28ee6b528f80600159c569a0e46ea81501d5f7d0a4161fbb67042952a73e500d01a5143272e6144b2d2498a0e34f3c910f2b943b94cb5e1642307264ad4317d456e5ebeb2acb06895d245b3429b4599dabfca351b561848f3223bb44c25aa56b58c0fa7893e8b542b963144f50512b3481adb72a87057cf9dabadbef8d9739f81cefde6351882bfc56a1d27e048dbbf9017f263a89ba921c38de0543c390d4750403d4958d7e55308adf03853d790c3bb09ca0362ecbe2e199a0bde104f3c346705cf24c3c13b2b39b20fe40310fe03d55c90bc19f5b426372202a550aae9e8bd59a3db5173f93ea774fce07c6f3bf92a27a6ea09ddd047159a2cc7e06b2373498757f9f6ed4165edda8f86c759ef02447391465ac4058801fcf1495acd496d0e21be5f74aca7204bcb212471856c851df09791b4192a612956a49c0c8a8be8f80269c62018b8ea93e33bd39b68b41bbf426a1b7b914154f833024c4e4099f13a683d0a4f7ebffa8be8f4f04c72024b59332de4a59d627b85800678559f514eaaa0c42e279ad343ede243c4793ebb9e04c0b19f896b82357d849852cfc10c91b54225e69315d8585a3d50bbc73359ada5d61ce9446b972ac177c48e82d3a0e446938303832ec184626156bcd4cbf53da34424bc5305f98b2c77fdac5e0f10e84bf58a5412d3c2f17ede171df06b3411ae5641e41ebac017a3104f23cb79f2e2fcfabeda3769f3560a1b0d12c13ce6880523a11f1bc3a245edd6e605355a3dc6f1fdc285b3a51498a67ab6e7f06c734ecf5e377a82fb51641185f5edf20d54e7fb1af7da315c1bacb6facfd0e1b55a3970909eb08746994a539217d033299e42d7818af5007ba8ccd358d9e68365e7b19751937f7d6d6c6e5de877e55f1b3d5292b3018e5d614e702ef46434430191d4c2b1f0f03a7c6d006f5682438dcccc36ebad8136f6b721f79dcf71ee7d0bf1bb77360bcee48e73f3199d8ca336eb2d8c68c93e1f94ed1383c70965bd664858521ee90b26c14ca626a706a6bacefc3d68267c2383a9a869d89bd7b9d8392bb088645ddc6fd9781d9d8badb90b3c6c3b48412bfc1a87a4f56e3a9689be5418e81a569c6cd7f26226802195ed696b3ccb8e612fa7fc712032a32fca4462d223279436f66636b080b161dadc5bfa7a90cc2f8c249cc8d69dbc954c94d5b414b041df12af2af78764c913fc98f0372ac89a5c5a68253ef2a1d54aa197fc3e992c1b5d37277c206c27df3ee060795c683ef08dc4d966e76b0c7bcb9e273f8f2d542b96f8b545b6992263ab148ed476d3b6c813c305134162a84dfe0b54d79b86a07f609b025bedd0fb231cbe0c52e5eaa96731d05ff10e13b1f63173d70dd317d467ba6dff5e4fe073ea8c65fd6463b101bb08c86ee6864a3108f8e6c7d32eaaa727404cdcfca78a93e9525ca40561c838c7e79fd356c6ee79277ad3c7be4c55f72dff57dde84b1e4f1db7bf363d6b5d947254fca12793ada1ed73187bbe05edd39c9687736e6d356936d75bc2a25e33a0bfcd9cbddece5ce8f26672ba3a1e38ca7b8ea7bd474b037f75be5799e61ad56e7d01e789ef1ab99684b640c1e95d12c35702c83ddae2f78631a1f274511f4400d0f1e9a9f5adadced463b319e78dc2c346d2b5e1dd62d23747dd36c17a9c93f010000ffff8638a060e7100000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("pb2", "./test")
		b.SetResolver("pb.interface.go", packr.Pointer{ForwardBox: gk, ForwardPath: "12a6b0a1bf0120c804d3cebd22ffe6e9"})
	}()
	return nil
}()
