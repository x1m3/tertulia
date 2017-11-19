package model

import "testing"
import (
	"time"
)

func TestTopicUUID(t *testing.T) {
	uuids := make(map[string]bool)
	for i := 0; i < 1000000; i++ {
		topic := NewTopic()
		if _, ok := uuids[topic.ID().String()]; ok {
			t.Error("We have a repeated UUID.. Bad thing")
		}
		uuids[topic.ID().String()] = true
	}
}

func TestTopicTitle(t *testing.T) {
	title := "El perro de San Roque No tiene Rabo"
	topic := NewTopic()
	topic.SetTitle(title)
	if topic.Title() != title {
		t.Errorf("Titles differ. Expecting <%s>, got <%s>", title, topic.Title())
	}
}

func TestTopicBody(t *testing.T) {
	body := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque maximus est vel consequat varius. Sed ut congue arcu. Maecenas facilisis nulla nec dapibus interdum. In ac nulla sollicitudin, suscipit lectus sit amet, congue ante. Nullam finibus turpis eget elementum laoreet. Mauris ac commodo nisi. Quisque varius ante nec odio mollis congue ac eget urna. Morbi consectetur massa sed sem congue ultrices. Mauris mattis lorem neque, at pretium mi fermentum sed. Quisque auctor sem ac ante ornare, ut congue nisl efficitur. In fringilla erat vel malesuada tempus. Cras purus urna, facilisis at purus ac, laoreet luctus dui. Sed nec tellus id enim consectetur efficitur eu a lorem. Sed sed nibh ac erat condimentum sollicitudin. Nam eu viverra turpis. Integer quis elit sit amet nisi suscipit venenatis. Maecenas mattis egestas nisi, at ultrices risus porta a. Integer aliquet, nibh vitae vestibulum vehicula, felis nunc ultricies nisi, vitae mollis ligula purus non mauris. Duis maximus dui a est sagittis iaculis. In vitae condimentum velit. Nunc vitae orci nibh. Nulla pretium leo quis sapien ultrices dictum. Proin tincidunt urna in enim elementum, vitae rhoncus justo volutpat. Ut pretium mattis nisi nec venenatis. Cras sit amet orci vitae risus placerat rhoncus. Aliquam at sapien ac metus auctor vehicula et vel sapien. Pellentesque sit amet justo dui. Morbi a ultricies tellus. Aenean et ante urna. Quisque mattis consequat tellus, vitae congue quam bibendum nec. Ut finibus porttitor sagittis. In hac habitasse platea dictumst. Phasellus eu posuere ligula, in rutrum dui. Duis pretium purus mauris, at tempus nibh ultricies vitae. In hac habitasse platea dictumst. Proin ac magna sapien. Cras efficitur ut sapien at volutpat. Vivamus sed justo laoreet, malesuada lacus id, fermentum ante. Etiam pretium nec elit vitae molestie. Nunc ligula nibh, porttitor et tincidunt sit amet, tincidunt nec tortor. Nullam fringilla ligula sed venenatis tempus. Integer vulputate nunc dui, sit amet vulputate lectus dictum at. Suspendisse rutrum, enim vel rhoncus rhoncus, nisl libero suscipit nibh, eget pulvinar tortor turpis in ligula. Duis eu est gravida urna tincidunt pharetra vitae ac sapien. Donec in quam facilisis, placerat nisl et, pellentesque sapien. Sed vestibulum est et lacus congue vulputate. In ut dolor ex. Etiam fringilla enim vel pulvinar facilisis. Nam ut tellus non augue interdum tempus id vel tellus. In nibh libero, sollicitudin in lacinia at, vulputate nec justo. Curabitur tincidunt, lacus et sollicitudin porttitor, tellus justo placerat enim, vitae posuere ligula nibh eget orci. Praesent euismod tristique sem, vulputate malesuada sem efficitur vel. Pellentesque non neque elementum, fringilla odio in, venenatis lectus. Sed mollis mollis enim ac commodo. Nulla mollis odio ac rhoncus tincidunt. Nam luctus accumsan mi nec dignissim. Aenean nec justo vel erat tristique congue vel eu justo. Vestibulum id euismod justo. Nullam pulvinar finibus euismod. Morbi elit tellus, gravida ac metus egestas, sodales venenatis dui. Nunc nec elementum libero. Cras condimentum tellus mauris, vitae euismod augue euismod ac. Maecenas tincidunt enim et erat interdum dapibus. Suspendisse euismod tempus arcu, vitae accumsan neque aliquet in. Donec tristique eros vel est hendrerit condimentum. In massa turpis, dapibus et leo at, sagittis condimentum ipsum. Vivamus et rhoncus urna. In rutrum urna a urna bibendum consequat. Vestibulum feugiat tempor ligula nec cursus. Pellentesque pulvinar quam in sagittis congue. Aliquam erat volutpat. Suspendisse faucibus odio a ipsum auctor, ut aliquam quam pretium. Pellentesque sed eleifend lacus. Maecenas ac felis massa. Suspendisse convallis nibh viverra augue interdum congue. Quisque imperdiet, nisi eget tincidunt aliquam, mauris nulla eleifend metus, a finibus ante turpis sit amet elit. Curabitur feugiat vel nisi at ornare. Integer rutrum risus dui, a sollicitudin ex porta sit amet. Praesent justo erat, lacinia vehicula eros in, rutrum vulputate ante. Fusce non nulla eu nibh maximus ornare. Nam mollis, dolor vitae sollicitudin interdum, nisl odio cursus massa, vel vehicula neque purus posuere sapien. Donec vel libero sit amet diam convallis bibendum a eget felis. Integer eleifend quam ut ligula imperdiet pellentesque. Quisque vitae gravida est, at faucibus justo. Aliquam sagittis ultricies fermentum. Praesent at libero suscipit, mattis mauris vel, cursus est. Donec commodo molestie risus, sit amet egestas lectus viverra vel. Fusce sollicitudin mattis ullamcorper. Sed luctus iaculis dui a suscipit. Integer eleifend sollicitudin leo, id placerat quam bibendum vel. Vestibulum ac condimentum diam. Integer a posuere ex. Cras hendrerit ut lacus vel ultricies. Etiam ut lectus est. Curabitur id libero luctus, mollis purus et, iaculis tortor. Duis fringilla purus vitae nibh placerat, in commodo urna aliquet. Sed rutrum gravida rutrum. Quisque a tempus augue. Aliquam sit amet urna faucibus, pharetra neque et, imperdiet lorem. Nullam quis sem non turpis pulvinar vehicula. Praesent eget rutrum diam. Suspendisse consequat et odio eu feugiat. Mauris eget neque in massa auctor commodo id sed sem. `
	topic := NewTopic()
	topic.SetBody(body)
	if topic.Body() != body {
		t.Errorf("Bodies differ. Expecting <%s>, got <%s>", body, topic.Body())
	}
}

func TestTopicTimes(t *testing.T) {
	topic := NewTopic()
	now := time.Now()
	if topic.CreationDate() != topic.ModDate() {
		t.Errorf("Creation and modification time differ. <%v> versus <%v>", topic.CreationDate(), topic.ModDate())
	}
	if now.Sub(topic.CreationDate()) > 1 * time.Millisecond {
		t.Errorf("Creation date is wrong (Or your server is a potato...)Expecting <%v>, got <%v>", now, topic.CreationDate())
	}
	oldCreationDate := topic.CreationDate()
	oldModDate := topic.ModDate()

	time.Sleep(1 * time.Millisecond)

	topic.SetTitle("San Roque's dog has no tail")
	topic.SetBody("Because Ramon Ramirez has cut it")
	if oldModDate == topic.ModDate() {
		t.Error("Updating a Topic is not  modifying modDate field")
	}
	if oldCreationDate != topic.CreationDate() {
		t.Error("Updating a Topic is changing creation date")
	}
}

func TestTopic_Author(t *testing.T) {
	topic := NewTopic()
	author:= NewPerson("manolito")
	topic.SetAuthor(author)
	if topic.Author().Nickname() != "manolito" {
		t.Error("Error assignign author to a topic")
	}
}

