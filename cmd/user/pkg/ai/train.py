import torch
import torchvision
from matplotlib import pyplot

device = torch.device("cpu")
class Model(torch.nn.Module):
    def __init__(self):
        super().__init__()
        self.transform = torchvision.transforms.Compose(
            [torchvision.transforms.ToTensor(),
             torchvision.transforms.Resize((224,224)),
             torchvision.transforms.Normalize((0.485,0.456,0.406),(0.229,0.224,0.225))
             ])
        self.net = torchvision.models.resnet50(pretrained=True).to(device).eval()

    def forward(self, image):
        image = pyplot.imread(image)
        image = self.transform(image)
        image = image.unsqueeze(0)
        image = self.net(image)
        return image

def extract(path):
    model = Model().to(device)
    result = model(path)
    result = result.squeeze(0)
    return str(result.tolist())
